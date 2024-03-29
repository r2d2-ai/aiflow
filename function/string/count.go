package string

import (
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	function.Register(&fnCount{})
}

type fnCount struct {
}

func (fnCount) Name() string {
	return "count"
}

func (fnCount) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnCount) Eval(params ...interface{}) (interface{}, error) {
	return strings.Count(params[0].(string), params[1].(string)), nil
}
