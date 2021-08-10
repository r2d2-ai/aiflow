package datetime

import (
	"time"

	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/support/log"

	"github.com/r2d2-ai/aiflow/core/data/expression/function"
)

type Now struct {
}

func init() {
	function.Register(&Now{})
}

func (s *Now) Name() string {
	return "now"
}

func (s *Now) GetCategory() string {
	return "datetime"
}

func (s *Now) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *Now) Eval(params ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with timezone")
	return time.Now().UTC().Format(DateTimeFormatDefault), nil
}
