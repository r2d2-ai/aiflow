package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/support/log"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type StartsWith struct {
}

func init() {
	function.Register(&StartsWith{})
}

func (s *StartsWith) Name() string {
	return "startsWith"
}

func (s *StartsWith) GetCategory() string {
	return "string"
}

func (s *StartsWith) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *StartsWith) Eval(params ...interface{}) (interface{}, error) {

	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.startsWith function first parameter [%+v] must be string", params[0])
	}

	substr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.startsWith function second parameter [%+v] must be string", params[1])
	}

	log.RootLogger().Debugf("Reports whether \"%s\" begins with \"%s\"", str, substr)

	return strings.HasPrefix(str, substr), nil
}
