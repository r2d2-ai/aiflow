package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"

	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

func init() {
	function.Register(&fnIndex{})
}

type fnIndex struct {
}

func (fnIndex) Name() string {
	return "index"
}

func (fnIndex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnIndex) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.index function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.index function second parameter [%+v] must be string", params[1])
	}
	return strings.Index(s1, s2), nil
}
