package datetime

import (
	"testing"

	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestCurrentDatetime_Eval(t *testing.T) {
	n := CurrentDatetime{}
	datetime, _ := n.Eval(nil)
	assert.NotNil(t, datetime)
}

func TestDatetime_CDT(t *testing.T) {
	n := CurrentDatetime{}
	date, _ := n.Eval(nil)
	assert.NotNil(t, date)
}
