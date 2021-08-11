package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	function.Register(&fnTrimLeft{})
}

type fnTrimLeft struct {
}

func (fnTrimLeft) Name() string {
	return "trimLeft"
}

func (fnTrimLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimLeft) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trimLeft function first parameter [%+v] must be string", params[0])
	}

	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.trimLeft function second parameter [%+v] must be string", params[1])
	}
	return strings.TrimLeft(s1, s2), nil
}
