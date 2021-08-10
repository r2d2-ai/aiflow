package datetime

import (
	"time"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/support/log"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
)

const DateTimeFormatDefault string = "2006-01-02T15:04:05-07:00"

type CurrentDatetime struct {
}

func init() {
	function.Register(&CurrentDatetime{})
}

func (s *CurrentDatetime) Name() string {
	return "currentDatetime"
}

func (s *CurrentDatetime) GetCategory() string {
	return "datetime"
}

func (s *CurrentDatetime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CurrentDatetime) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with timezone")
	return time.Now().UTC().Format(DateTimeFormatDefault), nil
}
