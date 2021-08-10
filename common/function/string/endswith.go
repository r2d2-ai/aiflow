package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"
	"github.com/r2d2-ai/AIflow/core/support/log"

	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

type EndsWith struct {
}

func init() {
	function.Register(&EndsWith{})
}

func (s *EndsWith) Name() string {
	return "endsWith"
}

func (s *EndsWith) GetCategory() string {
	return "string"
}

func (s *EndsWith) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *EndsWith) Eval(params ...interface{}) (interface{}, error) {

	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.contains function first parameter [%+v] must be string", params[0])
	}
	substr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.contains function second parameter [%+v] must be string", params[1])
	}
	log.RootLogger().Debugf("Reports whether \"%s\" ends with \"%s\"", str, substr)

	return strings.HasSuffix(str, substr), nil
}
