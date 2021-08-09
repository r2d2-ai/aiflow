package string

import (
	"testing"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnConcat_Eval(t *testing.T) {
	f := &fnConcat{}
	v, err := function.Eval(f, "a", "b")

	assert.Nil(t, err)
	assert.Equal(t, "ab", v)
}
