package string

import (
	"testing"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnContains_Eval(t *testing.T) {
	f := &fnContains{}

	v, err := function.Eval(f, "foo", "Bar")
	assert.Nil(t, err)
	assert.False(t, v.(bool))

	v, err = function.Eval(f, "foobar", "foo")
	assert.Nil(t, err)
	assert.True(t, v.(bool))
}
