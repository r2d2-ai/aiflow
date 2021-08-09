package string

import (
	"bytes"
	"fmt"

	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/data/coerce"
	"github.com/r2d2-ai/ai-box/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnConcat{})
}

type fnConcat struct {
}

func (fnConcat) Name() string {
	return "concat"
}

func (fnConcat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnConcat) Eval(params ...interface{}) (interface{}, error) {
	if len(params) >= 2 {
		var buffer bytes.Buffer

		for _, v := range params {
			s, err := coerce.ToString(v)
			if err != nil {
				return nil, fmt.Errorf("concat function parameter [%+v] must be string.", v)
			}
			buffer.WriteString(s)
		}
		return buffer.String(), nil
	}
	return "", fmt.Errorf("fnConcat function must have at least two arguments")
}
