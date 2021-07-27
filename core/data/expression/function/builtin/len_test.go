package builtin

import (
	"testing"

	"github.com/r2d2-ai/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnLen_Eval(t *testing.T) {
	f := &fnLen{}
	v, err := function.Eval(f, "abc")
	assert.Nil(t, err)
	assert.Equal(t, 3, v)

	v, err = function.Eval(f, []string{"a", "b"})
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
}
