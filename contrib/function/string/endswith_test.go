package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var end = &EndsWith{}

func TestStaticEndWith(t *testing.T) {
	final1, _ := end.Eval("TIBCO Project AIflow", "AIflow")
	assert.Equal(t, true, final1)

	final2, _ := end.Eval("TIBCO Web 集成器", "集成器")
	assert.Equal(t, true, final2)

	final3, _ := end.Eval("TIBCO 网路 AIflow", "网路")
	assert.Equal(t, false, final3)
}
