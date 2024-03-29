package datetime

import (
	"time"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type fnSub struct {
}

func init() {
	function.Register(&fnSub{})
}

func (s *fnSub) Name() string {
	return "sub"
}

func (s *fnSub) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt, data.TypeInt, data.TypeInt}, false
}

func (s *fnSub) Eval(in ...interface{}) (interface{}, error) {

	datetime, err := coerce.ToDateTime(in[0])
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

	year, month, day := datetime.Date()
	hour, min, sec := datetime.Clock()
	newDate := time.Date(year-years, month-time.Month(months), day-days, hour, min, sec, datetime.Nanosecond(), datetime.Location())
	return newDate, nil

}
