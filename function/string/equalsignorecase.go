package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	_ = function.Register(&fnEqualsIgnoreCase{})
}

type fnEqualsIgnoreCase struct {
}

func (s *fnEqualsIgnoreCase) Name() string {
	return "equalsIgnoreCase"
}

func (fnEqualsIgnoreCase) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnEqualsIgnoreCase) Eval(params ...interface{}) (interface{}, error) {
	str1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.contains function first parameter [%+v] must be string", params[0])
	}
	str2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.contains function second parameter [%+v] must be string", params[1])
	}
	return strings.EqualFold(str1, str2), nil
}
