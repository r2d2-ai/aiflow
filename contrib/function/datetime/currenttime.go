package datetime

import (
	"time"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/support/log"

	"github.com/r2d2-ai/ai-box/core/data/expression/function"
)

const TimeFormatDefault string = "15:04:05-07:00"

type CurrentTime struct {
}

func init() {
	function.Register(&CurrentTime{})
}

func (s *CurrentTime) Name() string {
	return "currentTime"
}

func (s *CurrentTime) GetCategory() string {
	return "datetime"
}

func (s *CurrentTime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CurrentTime) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current time with timezone")
	return time.Now().UTC().Format(TimeFormatDefault), nil
}
