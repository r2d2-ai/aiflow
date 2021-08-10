package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/data/coerce"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
)

func init() {
	function.Register(&fnToLower{})
}

type fnToLower struct {
}

func (fnToLower) Name() string {
	return "toLower"
}

func (fnToLower) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnToLower) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.toLower function first parameter [%+v] must be string", params[0])
	}
	return strings.ToLower(s1), nil
}
