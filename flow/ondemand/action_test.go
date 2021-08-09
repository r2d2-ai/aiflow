package ondemand

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/r2d2-ai/ai-box/core/support/test"

	"github.com/r2d2-ai/ai-box/core/action"
	_ "github.com/r2d2-ai/ai-box/core/support/test"

	"github.com/r2d2-ai/ai-box/core/engine/runner"
	"github.com/stretchr/testify/assert"
)

const testEventJson = `
{
  "payload": {
    "in1":"in1_value",
    "in2":"in2_value"
  },
  "AIflow" : {
      "inputs": { 
        "customerId": "=$.payload.in1",
          "orderId": "=$.payload.in2" 
        }
      ,
      "flow": {
        "metadata" : {
          "input":[
            { "name":"customerId", "type":"string" },
            { "name":"orderId", "type":"string" }
          ],
          "output":[
            { "name":"value", "type":"string" }
          ]
        },
        "tasks": [
          {
            "id": "testlog",
            "name": "testlog",
            "activity" : {
              "ref":"testlog",
              "input" : {
                "message" : "=$flow.orderId"
              }
            }
          }
        ]
      }
  }
}`

type Event struct {
	Payload interface{}     `json:"payload"`
	aiflow  json.RawMessage `json:"AIflow"`
}

//TestInitNoFlavorError
func TestFlowAction_Run(t *testing.T) {

	var evt Event

	// Unmarshall evt
	if err := json.Unmarshal([]byte(testEventJson), &evt); err != nil {
		assert.Nil(t, err)
		return
	}

	cfg := &action.Config{}

	ff := ActionFactory{}
	err := ff.Initialize(test.NewActionInitCtx())
	assert.Nil(t, err)

	fa, err := ff.New(cfg)
	assert.Nil(t, err)

	flowAction, ok := fa.(action.AsyncAction)
	assert.True(t, ok)

	inputs := make(map[string]interface{}, 2)

	inputs["flowPackage"] = evt.AIflow

	inputs["payload"] = evt.Payload

	r := runner.NewDirect()
	_, err = r.RunAction(context.Background(), flowAction, inputs)

	assert.Nil(t, err)
}
