package subflow

import (
	"context"
	"testing"

	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/engine/runner"

	"github.com/r2d2-ai/aiflow/core/action"
	"github.com/r2d2-ai/aiflow/core/activity"
	"github.com/r2d2-ai/aiflow/core/app/resource"
	"github.com/r2d2-ai/aiflow/core/support/test"
	"github.com/r2d2-ai/aiflow/flow"
	"github.com/r2d2-ai/aiflow/flow/support"
	"github.com/stretchr/testify/assert"
)

var jsonFlowRes1 = `{
"id":"flow:flow1",
"data":
{
  "tasks": [
    {
      "id": "runFlow",
      "activity": {
        "ref": "github.com/r2d2-ai/aiflow/flow/activity/subflow",
        "settings": {
          "flowURI": "res://flow/flow2"
        },
		"input": {
			"in":"test"
		}
      }
    }
  ]
}
}
`

var jsonFlowRes2 = `{
"id":"flow:flow2",
"data":
{
  "metadata": {
    "input": [
      {
        "name": "in",
        "type": "string"
      }
    ],
    "output": [
      {
        "name": "value",
        "type": "string"
      }
    ]
  },
  "tasks": [
    {
      "id": "log",
      "activity": {
        "ref": "testlog",
        "input": {
			"message":"=$flow.in"
        }
      }
    }
  ]
}
}`

var jsonFlow1 = `{
  "tasks": [
    {
      "id": "runFlow",
      "activity": {
        "ref": "github.com/r2d2-ai/aiflow/flow/activity/subflow",
        "settings": {
          "flowURI": "res://flow:flow2"
        },
		"input": {
			"in" : "test"
		}
      }
    }
  ]
}
`

var jsonFlow2 = `{
  "name":"the-subflow",
  "metadata": {
    "input": [
      {
        "name": "in",
        "type": "string"
      }
    ],
    "output": [
      {
        "name": "value",
        "type": "string"
      }
    ]
  },
  "tasks": [
    {
      "id": "log",
      "activity": {
        "ref": "testlog",
		"input" : {
			"message":"=$flow.in"
		}
      }
    }
  ]
}
`

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&SubFlowActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSettings(t *testing.T) {

	f := action.GetFactory("github.com/r2d2-ai/aiflow/flow")
	af := f.(*flow.ActionFactory)

	err := initActionFactory(af)
	assert.Nil(t, err)

	settings := &Settings{}
	iCtx := test.NewActivityInitContext(settings, nil)
	_, err = New(iCtx)
	assert.NotNil(t, err)

	settings = &Settings{FlowURI: "uri"}
	iCtx = test.NewActivityInitContext(settings, nil)
	a, err := New(iCtx)
	assert.NotNil(t, err) //invalid uri

	settings = &Settings{FlowURI: "res://flow:flow2"}
	iCtx = test.NewActivityInitContext(settings, nil)
	a, err = New(iCtx)
	assert.Nil(t, err)

	sfa, ok := a.(*SubFlowActivity)
	assert.True(t, ok)
	assert.Equal(t, "res://flow:flow2", sfa.flowURI)
}

func TestDynamicIO(t *testing.T) {

	f := action.GetFactory("github.com/r2d2-ai/aiflow/flow")
	af := f.(*flow.ActionFactory)

	err := initActionFactory(af)
	assert.Nil(t, err)

	settings := &Settings{FlowURI: "res://flow:flow2"}
	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	activityMd := act.Metadata()
	ioMd := activityMd.IOMetadata
	assert.NotNil(t, ioMd)
	tv, exists := ioMd.Input["in"]
	assert.True(t, exists)
	assert.Equal(t, data.TypeString, tv.Type())

	tv, exists = ioMd.Output["value"]
	assert.True(t, exists)
	assert.Equal(t, data.TypeString, tv.Type())
}

func TestSubFlow(t *testing.T) {

	f := action.GetFactory("github.com/r2d2-ai/aiflow/flow")
	af := f.(*flow.ActionFactory)

	err := initActionFactory(af)
	assert.Nil(t, err)

	flowAction, err := f.New(&action.Config{Settings: map[string]interface{}{"flowURI": "res://flow:flow1"}})

	assert.Nil(t, err)
	assert.NotNil(t, flowAction)

	dr := runner.NewDirect()
	results, err := dr.RunAction(context.Background(), flowAction, nil)
	assert.Nil(t, err)
	assert.Nil(t, results)
}

func initActionFactory(af action.Factory) error {

	ctx := test.NewActionInitCtx()
	err := af.Initialize(ctx)
	if err != nil {
		return err
	}

	rConfig1 := &resource.Config{ID: "flow:flow1", Data: []byte(jsonFlow1)}
	rConfig2 := &resource.Config{ID: "flow:flow2", Data: []byte(jsonFlow2)}

	err = ctx.AddResource(support.ResTypeFlow, rConfig1)
	if err != nil {
		return err
	}

	err = ctx.AddResource(support.ResTypeFlow, rConfig2)
	if err != nil {
		return err
	}

	return err
}
