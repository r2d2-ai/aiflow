package stream

import (
	"encoding/json"
	"testing"

	"github.com/r2d2-ai/aiflow/action"
	"github.com/r2d2-ai/aiflow/action/stream/pipeline"
	"github.com/r2d2-ai/aiflow/app/resource"
	"github.com/r2d2-ai/aiflow/engine/channels"
	"github.com/r2d2-ai/aiflow/support/test"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
  "id": "AIflow-stream",
  "ref": "github.com/r2d2-ai/aiflow/action/stream",
  "settings": {
    "streamURI": "res://stream:test",
    "outputChannel": "testChan"
  }
}
`
const resData string = `{
        "metadata": {
          "input": [
            {
              "name": "input",
              "type": "integer"
            }
          ]
        },
        "stages": [
        ]
      }`

func TestActionFactory_New(t *testing.T) {

	cfg := &action.Config{}
	err := json.Unmarshal([]byte(testConfig), cfg)

	if err != nil {
		t.Error(err)
		return
	}

	af := ActionFactory{}
	ctx := test.NewActionInitCtx()

	err = af.Initialize(ctx)
	assert.Nil(t, err)

	resourceCfg := &resource.Config{ID: "stream:test"}
	resourceCfg.Data = []byte(resData)
	err = ctx.AddResource(pipeline.ResType, resourceCfg)
	assert.Nil(t, err)

	_, err = channels.New("testChan", 5)
	assert.Nil(t, err)

	act, err := af.New(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, act)
}
