package datetime

import (
	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type fnAdd struct {
}

func init() {
	function.Register(&fnAdd{})
}

func (s *fnAdd) Name() string {
	return "add"
}

func (s *fnAdd) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt, data.TypeInt, data.TypeInt}, false
}

func (s *fnAdd) Eval(in ...interface{}) (interface{}, error) {
	startDate, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	years, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	months, err := coerce.ToInt(in[2])
	if err != nil {
		return nil, err
	}
	days, err := coerce.ToInt(in[3])
	if err != nil {
		return nil, err
	}
	newT := startDate.AddDate(years, months, days)
	return newT, nil

}
