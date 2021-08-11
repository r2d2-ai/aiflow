package engine

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/r2d2-ai/aiflow/app"
	"github.com/r2d2-ai/aiflow/data/schema"
	"github.com/r2d2-ai/aiflow/engine/secret"
	"github.com/r2d2-ai/aiflow/support"
)

var appName, appVersion string

func init() {
	if IsSchemaSupportEnabled() {
		schema.Enable()

		if !IsSchemaValidationEnabled() {
			schema.DisableValidation()
		}
	}
}

// Returns name of the application
func GetAppName() string {
	return appName
}

// Returns version of the application
func GetAppVersion() string {
	return appVersion
}

func LoadAppConfig(flowJson string, compressed bool) (*app.Config, error) {

	var jsonBytes []byte

	if flowJson == "" {

		// a json string wasn't provided, so lets lookup the file in path
		configPath := GetAIflowAppConfigPath()

		flow, err := os.Open(configPath)
		if err != nil {
			return nil, err
		}

		jsonBytes, err = ioutil.ReadAll(flow)
		if err != nil {
			return nil, err
		}
	} else {

		if compressed {
			var err error
			jsonBytes, err = support.DecodeAndUnzip(flowJson)
			if err != nil {
				return nil, err
			}
		} else {
			jsonBytes = []byte(flowJson)
		}
	}

	updated, err := secret.PreProcessConfig(jsonBytes)
	if err != nil {
		return nil, err
	}

	appConfig := &app.Config{}
	err = json.Unmarshal(updated, &appConfig)
	if err != nil {
		return nil, err
	}

	appName = appConfig.Name
	appVersion = appConfig.Version

	return appConfig, nil
}
