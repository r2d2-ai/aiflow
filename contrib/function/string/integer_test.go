package string

import (
	"testing"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnInteger_Eval(t *testing.T) {
	f := &fnInteger{}
	v, err := function.Eval(f, "123")

	assert.Nil(t, err)
	assert.Equal(t, 123, v)
}
