package string

import (
	"testing"

	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnFloat_Eval(t *testing.T) {
	f := &fnFloat{}
	v, err := function.Eval(f, "123.123")

	assert.Nil(t, err)
	assert.Equal(t, 123.123, v)
}
