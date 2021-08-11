package timer

import (
	"encoding/json"
	"testing"

	"github.com/r2d2-ai/aiflow/action"
	"github.com/r2d2-ai/aiflow/support/test"
	"github.com/r2d2-ai/aiflow/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "AIflow-timer",
	"ref": "github.com/r2d2-ai/aiflow/common/trigger/timer",
	"handlers": [
	  {
		"settings":{
			"repeatInterval" : "1s"
		},
		"action":{
			"id":"dummy"
		}
	  }
	]
  }
  `

func TestInitOk(t *testing.T) {
	f := &Factory{}
	tgr, err := f.New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, tgr)
}

func TestTimerTrigger_Initialize(t *testing.T) {
	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)

}
