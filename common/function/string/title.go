package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/AIflow/core/data"
	"github.com/r2d2-ai/AIflow/core/data/coerce"
	"github.com/r2d2-ai/AIflow/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTitle{})
}

type fnTitle struct {
}

func (fnTitle) Name() string {
	return "toTitleCase"
}

func (fnTitle) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnTitle) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("error converting string.toTitleCase's argument [%+v] to string: %s", params[0], err.Error())
	}
	return strings.Title(str), nil
}
