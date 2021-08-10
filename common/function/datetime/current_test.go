package datetime

import (
	"fmt"
	"testing"

	"github.com/r2d2-ai/AIflow/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentDate_Eval(t *testing.T) {
	n := currentFn{}
	datetime, _ := n.Eval(nil)
	fmt.Println(datetime)
	assert.NotNil(t, datetime)
}
