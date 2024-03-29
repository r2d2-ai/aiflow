package datetime

import (
	"fmt"
	"time"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type fnParseTIme struct {
}

func init() {
	function.Register(&fnParseTIme{})
}

func (s *fnParseTIme) Name() string {
	return "parse"
}

func (s *fnParseTIme) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, true
}

func (s *fnParseTIme) Eval(params ...interface{}) (interface{}, error) {

	parsedTime, err := coerce.ToDateTime(params[0])
	if err != nil {
		return nil, err
	}
	if len(params) >= 2 {
		zone, err := coerce.ToString(params[1])
		if err != nil {
			return nil, fmt.Errorf("Format date second argument must be string")
		}
		if len(zone) <= 0 {
			zone = "UTC"
		}
		loc, err := time.LoadLocation(zone)
		if err != nil {
			return nil, err
		}
		newTime := parsedTime.In(loc)
		parsedTime = newTime

	}

	return parsedTime, nil
}
