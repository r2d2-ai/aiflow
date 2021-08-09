package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/data/coerce"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
)

func init() {
	function.Register(&fnIndexAny{})
}

type fnIndexAny struct {
}

func (fnIndexAny) Name() string {
	return "indexAny"
}

func (fnIndexAny) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnIndexAny) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.indexAny function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.indexAny function second parameter [%+v] must be string", params[1])
	}
	return strings.IndexAny(s1, s2), nil
}
