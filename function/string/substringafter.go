package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/support/log"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type Substringafter struct {
}

func init() {
	function.Register(&Substringafter{})
}

func (s *Substringafter) Name() string {
	return "substringAfter"
}

func (s *Substringafter) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *Substringafter) Eval(params ...interface{}) (interface{}, error) {

	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.substringAfter function first parameter [%+v] must be string", params[0])
	}

	afterStr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.substringAfter function second parameter [%+v] must be integer", params[1])
	}

	log.RootLogger().Debugf("Start Substringafter function with parameters string %s after string %s", str, afterStr)
	if strings.Index(str, afterStr) >= 0 {
		return str[strings.Index(str, afterStr)+len(afterStr):], nil
	} else {
		return str, nil
	}
}
