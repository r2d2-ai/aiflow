package array

import (
	"fmt"
	"reflect"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/expression/function"
	"github.com/r2d2-ai/AIflow/core/support/log"
)

type mergeFunc struct {
}

func init() {
	function.Register(&mergeFunc{})
}

func (a *mergeFunc) Name() string {
	return "merge"
}

func (mergeFunc) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, true
}

func (mergeFunc) Eval(params ...interface{}) (interface{}, error) {

	if len(params) < 2 {
		return nil, fmt.Errorf("array merge must have at least 2 array")
	}

	finalArrayValue := reflect.Value{}
	log.RootLogger().Trace("Start array mergeFunc function with parameters %+v", params)
	for _, arg := range params {
		if arg != nil {
			arrV := reflect.ValueOf(arg)
			if arrV.Kind() == reflect.Slice {
				if !finalArrayValue.IsValid() {
					finalArrayValue = arrV
					continue
				} else {
					if arrV.Kind() == reflect.Slice {
						for i := 0; i < arrV.Len(); i++ {
							finalArrayValue = reflect.Append(finalArrayValue, arrV.Index(i))
						}
					} else {
						finalArrayValue = reflect.Append(finalArrayValue, arrV)
					}
				}

			}
		}
	}

	log.RootLogger().Trace("array append function done, final array %+v", finalArrayValue.Interface())
	return finalArrayValue.Interface(), nil
}
