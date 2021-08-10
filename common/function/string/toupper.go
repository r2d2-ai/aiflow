package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"

	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

func init() {
	function.Register(&fnToUpper{})
}

type fnToUpper struct {
}

func (fnToUpper) Name() string {
	return "toUpper"
}

func (fnToUpper) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnToUpper) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.toUpper function first parameter [%+v] must be string", params[0])
	}
	return strings.ToUpper(s1), nil
}
