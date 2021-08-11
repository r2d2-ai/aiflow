package datetime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type fnSubMins struct {
}

func init() {
	function.Register(&fnSubMins{})
}

func (s *fnSubMins) Name() string {
	return "subMins"
}

func (s *fnSubMins) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeInt}, false
}

func (s *fnSubMins) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	mins, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration("-" + strconv.Itoa(mins) + "m")
	if err != nil {
		return nil, fmt.Errorf("Invalid minutes [%d]", mins)
	}
	return t.Add(d), nil

}
