package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/support/log"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type Substringbefore struct {
}

func init() {
	function.Register(&Substringbefore{})
}

func (s *Substringbefore) Name() string {
	return "substringBefore"
}

func (s *Substringbefore) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *Substringbefore) Eval(params ...interface{}) (interface{}, error) {

	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.substringAfter function first parameter [%+v] must be string", params[0])
	}

	beforeStr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.substringAfter function second parameter [%+v] must be integer", params[1])
	}

	log.RootLogger().Debugf("Start substringbefore function with parameters string %s before string %s", str, beforeStr)
	if strings.Index(str, beforeStr) >= 0 {
		return str[:strings.Index(str, beforeStr)], nil
	} else {
		return str, nil
	}
}
