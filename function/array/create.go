package array

import (
	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/r2d2-ai/aiflow/support/log"
)

type Create struct {
}

func init() {
	function.Register(&Create{})
}

func (s *Create) Name() string {
	return "create"
}

func (s *Create) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, true
}

func (s *Create) Eval(object ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Start array function with parameters %v", object)
	if object == nil {
		return nil, nil
	}
	var result []interface{}
	result = append(result, object...)
	log.RootLogger().Debugf("Done array function with result %v", result)
	return result, nil
}
