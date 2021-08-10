package datetime

import (
	"time"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"
	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

type fnAddSeconds struct {
}

func init() {
	function.Register(&fnAddSeconds{})
}

func (s *fnAddSeconds) Name() string {
	return "addSeconds"
}

func (s *fnAddSeconds) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnAddSeconds) Eval(in ...interface{}) (interface{}, error) {

	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}

	seconds, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	return t.Add(time.Duration(seconds) * time.Second), nil
}
