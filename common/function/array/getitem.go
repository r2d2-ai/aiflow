package array

import (
	"fmt"
	"reflect"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/data/coerce"
	"github.com/r2d2-ai/ai-box/core/data/expression/function"
	"github.com/r2d2-ai/ai-box/core/support/log"
)

type Get struct {
}

func init() {
	function.Register(&Get{})
}

func (s *Get) Name() string {
	return "get"
}

func (s *Get) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeInt}, false
}

func (s *Get) Eval(params ...interface{}) (interface{}, error) {

	items := params[0]
	index, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("array get function second parameter must be integer")
	}

	log.RootLogger().Debugf("Start array get function with parameters %+v and %+v", items, index)

	if items == nil {
		return nil, fmt.Errorf("index out of bounds, index [%d] but array empty", index)
	}

	arrV := reflect.ValueOf(items)
	if arrV.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unable to use array.get on un-array object")
	}

	length := arrV.Len()
	if index >= length {
		return nil, fmt.Errorf("index out of bounds, index [%d] but array length [%d]", index, length)
	} else {
		return arrV.Index(index).Interface(), nil
	}
}
