import com.sap.piper.BashUtils
import com.sap.piper.DefaultValueCache
import com.sap.piper.JenkinsUtils
import com.sap.piper.MapUtils
import com.sap.piper.PiperGoUtils
import com.sap.piper.Utils

import static com.sap.piper.Prerequisites.checkScript

void call(Map parameters = [:], stepName, metadataFile, List credentialInfo, failOnMissingReports = false, failOnMissingLinks = false) {

    handlePipelineStepErrors(stepName: stepName, stepParameters: parameters) {

        def stepParameters = [:].plus(parameters)

        def script = checkScript(this, parameters) ?: this
        stepParameters.remove('script')

        def utils = parameters.juStabUtils ?: new Utils()
        stepParameters.remove('juStabUtils')

        def jenkinsUtils = parameters.jenkinsUtilsStub ?: new JenkinsUtils()
        stepParameters.remove('jenkinsUtilsStub')

        new PiperGoUtils(this, utils).unstashPiperBin()
        utils.unstash('pipelineConfigAndTests')
        script.commonPipelineEnvironment.writeToDisk(script)

        writeFile(file: ".pipeline/tmp/${metadataFile}", text: libraryResource(metadataFile))

        // When converting to JSON and back again, entries which had a 'null' value will now have a value
        // of type 'net.sf.json.JSONNull', for which the Groovy Truth resolves to 'true' in for example if-conditions
        stepParameters = MapUtils.pruneNulls(stepParameters)

        withEnv([
            "PIPER_parametersJSON=${groovy.json.JsonOutput.toJson(stepParameters)}",
            //ToDo: check if parameters make it into docker image on JaaS
        ]) {
            String defaultConfigArgs = getCustomDefaultConfigsArg()
            String customConfigArg = getCustomConfigArg(script)

            // get context configuration
            Map config = readJSON(text: sh(returnStdout: true, script: "./piper getConfig --contextConfig --stepMetadata '.pipeline/tmp/${metadataFile}'${defaultConfigArgs}${customConfigArg}"))
            echo "Config: ${config}"

            dockerWrapper(script, config) {
                credentialWrapper(config, credentialInfo) {
                    sh 'ls -a /'
                    int stat = sh(script: 'ls -a s4hana_pipeline/maven_local_repo/', returnStatus: true)
                    int status = sh(script: 'ls -a /var/jenkins_home/', returnStatus: true)
                    sh "printenv | sort"
                    sh "./piper ${stepName}${defaultConfigArgs}${customConfigArg}"
                    //sh 'ls -a /root/.m2'
                }
                jenkinsUtils.handleStepResults(stepName, failOnMissingReports, failOnMissingLinks)
                script.commonPipelineEnvironment.readFromDisk(script)
            }
        }
    }
}

static String getCustomDefaultConfigs() {
    // The default config files were extracted from merged library
    // resources by setupCommonPipelineEnvironment.groovy into .pipeline/.
    List customDefaults = DefaultValueCache.getInstance().getCustomDefaults()
    for (int i = 0; i < customDefaults.size(); i++) {
        customDefaults[i] = BashUtils.quoteAndEscape(".pipeline/${customDefaults[i]}")
    }
    return customDefaults.join(',')
}

static String getCustomDefaultConfigsArg() {
    String customDefaults = getCustomDefaultConfigs()
    if (customDefaults) {
        return " --defaultConfig ${customDefaults}"
    }
    return ''
}

static String getCustomConfigArg(def script) {
    if (script?.commonPipelineEnvironment?.configurationFile
        && script.commonPipelineEnvironment.configurationFile != '.pipeline/config.yml'
        && script.commonPipelineEnvironment.configurationFile != '.pipeline/config.yaml') {
        return " --customConfig ${BashUtils.quoteAndEscape(script.commonPipelineEnvironment.configurationFile)}"
    }
    return ''
}

void dockerWrapper(script, config, body) {
    if (config.dockerImage) {
        dockerExecute(
            script: script,
            dockerImage: config.dockerImage,
            dockerWorkspace: config.dockerWorkspace,
            dockerOptions: config.dockerOptions,
            //ToDo: add additional dockerExecute parameters
        ) {
            body()
        }
    } else {
        body()
    }
}

void credentialWrapper(config, List credentialInfo, body) {
    if (credentialInfo.size() > 0) {
        def creds = []
        def sshCreds = []
        credentialInfo.each { cred ->
            switch(cred.type) {
                case "file":
                    if (config[cred.id]) creds.add(file(credentialsId: config[cred.id], variable: cred.env[0]))
                    break
                case "token":
                    if (config[cred.id]) creds.add(string(credentialsId: config[cred.id], variable: cred.env[0]))
                    break
                case "usernamePassword":
                    if (config[cred.id]) creds.add(usernamePassword(credentialsId: config[cred.id], usernameVariable: cred.env[0], passwordVariable: cred.env[1]))
                    break
                case "ssh":
                    if (config[cred.id]) sshCreds.add(config[cred.id])
                    break
                default:
                    error ("invalid credential type: ${cred.type}")
            }
        }

        if (sshCreds.size() > 0) {
            sshagent (sshCreds) {
                withCredentials(creds) {
                    body()
                }
            }
        } else {
            withCredentials(creds) {
                body()
            }
        }
    } else {
        body()
    }
}
