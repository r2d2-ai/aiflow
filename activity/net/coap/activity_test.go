package coap

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

func TestSimple(t *testing.T) {
	settings := &Settings{Method: "GET", URI: "coap://localhost:5683/AIflow"}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	act.Eval(tc)
}
