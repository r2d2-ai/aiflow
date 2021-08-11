package datetime

import (
	"time"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/support/log"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

type currentFn struct {
}

func init() {
	function.Register(&currentFn{})
}

func (s *currentFn) Name() string {
	return "current"
}

func (s *currentFn) GetCategory() string {
	return "datetime"
}

func (s *currentFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *currentFn) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with UTC timezone")
	return time.Now().UTC(), nil
}
