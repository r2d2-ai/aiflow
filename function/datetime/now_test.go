package datetime

import (
	"fmt"
	"testing"

	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestNow_Eval(t *testing.T) {
	n := Now{}
	now, _ := n.Eval(nil)
	assert.NotNil(t, now)
	fmt.Println(now)
}
