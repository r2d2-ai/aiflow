package engine

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEngineConfig = `
{
  "type": "AIflow:engine",
  "imports": [
    "github.com/r2d2-ai/aiflow/service/flow-state/store/mem"
  ],
  "actionSettings": {
    "github.com/r2d2-ai/aiflow/action/flow": {
      "stepRecordingMode": "full",
      "snapshotRecordingMode": "off",
      "enableExternalFlows": true
    }
  },
  "services": [
    {
      "name": "flowTester",
      "ref": "github.com/r2d2-ai/aiflow/action/flow/tester",
      "enabled": true,
      "settings": {
        "port": "8181"
      }
    },
    {
      "name": "flowStateRecorder",
      "ref": "github.com/r2d2-ai/aiflow/service/flow-state/client/local",
      "enabled": true,
      "settings": {
      }
    },
    {
      "name": "flowStateProvider",
      "ref": "github.com/r2d2-ai/aiflow/service/flow-state/server/rest",
      "enabled": true,
      "settings": {
        "host": "blah",
        "port": "8080"
      }
    }
  ]
}
`

func TestConfigUnmarshal(t *testing.T) {

	config := &Config{}
	err := json.Unmarshal([]byte(testEngineConfig), config)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(config.Imports))
	assert.Equal(t, 1, len(config.ActionSettings))
	assert.Equal(t, 3, len(config.Services))
}

func TestLoadEngineConfig(t *testing.T) {

	config, err := LoadEngineConfig(testEngineConfig, false)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(config.Imports))
	assert.Equal(t, 1, len(config.ActionSettings))
	assert.Equal(t, 3, len(config.Services))
	assert.Equal(t, "POOLED", config.RunnerType)
	assert.True(t, config.StopEngineOnError)
}
