// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type cloudFoundryCreateServiceKeyOptions struct {
	CfAPIEndpoint      string `json:"cfApiEndpoint,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	CfOrg              string `json:"cfOrg,omitempty"`
	CfSpace            string `json:"cfSpace,omitempty"`
	CfServiceInstance  string `json:"cfServiceInstance,omitempty"`
	CfServiceKeyName   string `json:"cfServiceKeyName,omitempty"`
	CfServiceKeyConfig string `json:"cfServiceKeyConfig,omitempty"`
}

// CloudFoundryCreateServiceKeyCommand cloudFoundryCreateServiceKey
func CloudFoundryCreateServiceKeyCommand() *cobra.Command {
	metadata := cloudFoundryCreateServiceKeyMetadata()
	var stepConfig cloudFoundryCreateServiceKeyOptions
	var startTime time.Time

	var createCloudFoundryCreateServiceKeyCmd = &cobra.Command{
		Use:   "cloudFoundryCreateServiceKey",
		Short: "cloudFoundryCreateServiceKey",
		Long:  `Create CloudFoundryServiceKey`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("cloudFoundryCreateServiceKey")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "cloudFoundryCreateServiceKey", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "cloudFoundryCreateServiceKey")
			cloudFoundryCreateServiceKey(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addCloudFoundryCreateServiceKeyFlags(createCloudFoundryCreateServiceKeyCmd, &stepConfig)
	return createCloudFoundryCreateServiceKeyCmd
}

func addCloudFoundryCreateServiceKeyFlags(cmd *cobra.Command, stepConfig *cloudFoundryCreateServiceKeyOptions) {
	cmd.Flags().StringVar(&stepConfig.CfAPIEndpoint, "cfApiEndpoint", os.Getenv("PIPER_cfApiEndpoint"), "Cloud Foundry API endpoint")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User or E-Mail for CF")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "User Password for CF User")
	cmd.Flags().StringVar(&stepConfig.CfOrg, "cfOrg", os.Getenv("PIPER_cfOrg"), "CF org")
	cmd.Flags().StringVar(&stepConfig.CfSpace, "cfSpace", os.Getenv("PIPER_cfSpace"), "CF Space")
	cmd.Flags().StringVar(&stepConfig.CfServiceInstance, "cfServiceInstance", os.Getenv("PIPER_cfServiceInstance"), "Parameter of ServiceInstance Name to delete CloudFoundry Service")
	cmd.Flags().StringVar(&stepConfig.CfServiceKeyName, "cfServiceKeyName", os.Getenv("PIPER_cfServiceKeyName"), "Parameter of CloudFoundry Service Key to be created")
	cmd.Flags().StringVar(&stepConfig.CfServiceKeyConfig, "cfServiceKeyConfig", os.Getenv("PIPER_cfServiceKeyConfig"), "Configuration for Service Key Creation")

	cmd.MarkFlagRequired("cfApiEndpoint")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("cfOrg")
	cmd.MarkFlagRequired("cfSpace")
	cmd.MarkFlagRequired("cfServiceInstance")
	cmd.MarkFlagRequired("cfServiceKeyName")
}

// retrieve step metadata
func cloudFoundryCreateServiceKeyMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "cloudFoundryCreateServiceKey",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "cfApiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "cloudFoundry/apiEndpoint"}},
					},
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cfOrg",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "cloudFoundry/org"}},
					},
					{
						Name:        "cfSpace",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "cloudFoundry/space"}},
					},
					{
						Name:        "cfServiceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceInstance"}},
					},
					{
						Name:        "cfServiceKeyName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceKeyName"}},
					},
					{
						Name:        "cfServiceKeyConfig",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceKeyConfig"}},
					},
				},
			},
		},
	}
	return theMetaData
}
