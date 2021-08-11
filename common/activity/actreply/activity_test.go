package actreply

import (
	"encoding/json"
	"testing"

	"github.com/r2d2-ai/aiflow/core/activity"
	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/data/mapper"
	"github.com/r2d2-ai/aiflow/core/data/metadata"
	"github.com/r2d2-ai/aiflow/core/data/resolve"
	"github.com/r2d2-ai/aiflow/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSimpleReply(t *testing.T) {

	//set mappings
	mappingsJson :=
		`{
           "Output1": "1",
           "Output2": 2
		}`

	var mappings interface{}
	err := json.Unmarshal([]byte(mappingsJson), &mappings)
	if err != nil {
		panic("Unable to parse mappings: " + err.Error())
	}

	settings := map[string]interface{}{"mappings": mappings}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	ac := newActionContext()
	tc := test.NewActivityContextWithAction(act.Metadata(), ac)

	//eval
	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)

	assert.Nil(t, ac.ReplyErr)
	o1, exists1 := ac.ReplyData["Output1"]
	assert.True(t, exists1, "Output1 not set")
	if exists1 {
		assert.Equal(t, "1", o1)
	}
	o2, exists2 := ac.ReplyData["Output2"]
	assert.True(t, exists2, "Output2 not set")
	if exists2 {
		assert.Equal(t, 2.0, o2)
	}
}

func newActionContext() *test.TestActivityHost {
	input := map[string]data.TypedValue{"Input1": data.NewTypedValue(data.TypeString, "")}
	output := map[string]data.TypedValue{"Output1": data.NewTypedValue(data.TypeString, ""), "Output2": data.NewTypedValue(data.TypeInt, "")}

	ac := &test.TestActivityHost{
		HostId:     "1",
		HostRef:    "github.com/r2d2-ai/aiflow/action/flow",
		IoMetadata: &metadata.IOMetadata{Input: input, Output: output},
		HostData:   data.NewSimpleScope(nil, nil),
	}

	return ac
}
