package datetime

import (
	"time"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"
	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

type fnAddHours struct {
}

func init() {
	function.Register(&fnAddHours{})
}

func (s *fnAddHours) Name() string {
	return "addHours"
}

func (s *fnAddHours) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnAddHours) Eval(in ...interface{}) (interface{}, error) {
	startDate, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	hours, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	newT := startDate.Add(time.Duration(hours) * time.Hour)
	return newT, nil
}
