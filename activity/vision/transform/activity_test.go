package transform

import (
	"testing"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	act.Eval(tc)
}

func TestAddToFlow(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	act.Eval(tc)
}
