package coap

import (
	"encoding/json"
	"testing"

	"github.com/r2d2-ai/aiflow/action"
	"github.com/r2d2-ai/aiflow/support"
	"github.com/r2d2-ai/aiflow/support/test"
	"github.com/r2d2-ai/aiflow/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}

const testConfig string = `{
	"id": "trigger-coap",
	"ref": "github.com/r2d2-ai/aiflow/device-contrib/trigger/coap",
	"handlers": [
	  {
			"settings": {
		  	"path": "/AIflow"
			},
			"action" : {
		  	"id": "dummy"
			}
		}
	]
  }`

func TestRestTrigger_Initialize(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
	})}

	trg, err := test.InitTrigger(f, config, actions)

	assert.Nil(t, err)
	assert.NotNil(t, trg)
	err = trg.Start()

	assert.Nil(t, err)
}
