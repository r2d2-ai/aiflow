package datetime

import (
	"time"

	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/support/log"

	"github.com/r2d2-ai/aiflow/core/data/expression/function"
)

const DateFormatDefault = "2006-01-02-07:00"

type CurrentDate struct {
}

func init() {
	function.Register(&CurrentDate{})
}

func (s *CurrentDate) Name() string {
	return "currentDate"
}

func (s *CurrentDate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CurrentDate) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current date with timezone")
	return time.Now().UTC().Format(DateFormatDefault), nil
}
