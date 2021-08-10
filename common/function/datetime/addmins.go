package datetime

import (
	"time"

	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/data/coerce"
	"github.com/r2d2-ai/aiflow/core/data/expression/function"
)

type fnAddMins struct {
}

func init() {
	function.Register(&fnAddMins{})
}

func (s *fnAddMins) Name() string {
	return "addMins"
}

func (s *fnAddMins) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnAddMins) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	mins, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	return t.Add(time.Duration(mins) * time.Minute), nil

}
