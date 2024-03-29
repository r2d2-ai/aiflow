package array

import (
	"fmt"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/r2d2-ai/aiflow/support/log"

	"reflect"
)

type Count struct {
}

func init() {
	function.Register(&Count{})
}

func (s *Count) Name() string {
	return "count"
}

func (s *Count) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (s *Count) Eval(params ...interface{}) (interface{}, error) {
	items := params[0]
	log.RootLogger().Debugf("Start array count function with parameters %v", items)
	if items == nil {
		return 0, nil
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() == reflect.Slice {
		return arrV.Len(), nil
	}

	return 0, fmt.Errorf("unable to count un-array object")
}
