package string

import (
	"fmt"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	_ = function.Register(&fnLen{})
}

type fnLen struct {
}

func (fnLen) Name() string {
	return "len"
}

func (fnLen) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnLen) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.len function first parameter [%+v] must be string", params[0])
	}

	return len(s1), nil
}
