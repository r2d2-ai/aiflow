package datetime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type fnSubHours struct {
}

func init() {
	function.Register(&fnSubHours{})
}

func (s *fnSubHours) Name() string {
	return "subHours"
}

func (s *fnSubHours) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnSubHours) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	hours, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	d, err := time.ParseDuration("-" + strconv.Itoa(hours) + "h")
	if err != nil {
		return nil, fmt.Errorf("Invalid hours [%d]", hours)
	}
	return t.Add(d), nil
}
