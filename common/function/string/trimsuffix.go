package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/data/coerce"

	"github.com/r2d2-ai/aiflow/core/data/expression/function"
)

func init() {
	function.Register(&fnTrimSuffix{})
}

type fnTrimSuffix struct {
}

func (fnTrimSuffix) Name() string {
	return "trimSuffix"
}

func (fnTrimSuffix) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnTrimSuffix) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.trimSuffix function first parameter [%+v] must be string", params[0])
	}

	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.trimSuffix function second parameter [%+v] must be string", params[1])
	}

	return strings.TrimSuffix(s1, s2), nil
}
