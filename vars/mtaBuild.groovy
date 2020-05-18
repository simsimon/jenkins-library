import com.sap.piper.DownloadCacheUtils
import static com.sap.piper.Prerequisites.checkScript
import groovy.transform.Field

@Field String STEP_NAME = getClass().getName()
//@Field String METADATA_FILE = 'metadata/mtaBuild.yaml'

void call(Map parameters = [:]) {
    final script = checkScript(this, parameters) ?: this
    parameters = DownloadCacheUtils.injectDownloadCacheInMavenParameters(script, parameters)

    piperExecuteBin parameters, STEP_NAME, "metadata/${STEP_NAME}.yaml", []
}

