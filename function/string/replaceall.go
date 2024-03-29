package string

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"

	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	function.Register(&fnReplaceAll{})
}

type fnReplaceAll struct {
}

func (fnReplaceAll) Name() string {
	return "replaceAll"
}

func (fnReplaceAll) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (fnReplaceAll) Eval(params ...interface{}) (interface{}, error) {

	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.replaceAll function first parameter [%+v] must be string", params[0])
	}
	s2, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.replaceAll function second parameter [%+v] must be string", params[1])
	}
	s3, err := coerce.ToString(params[2])
	if err != nil {
		return nil, fmt.Errorf("string.replaceAll function third parameter [%+v] must be string", params[2])
	}
	return strings.Replace(s1, s2, s3, -1), nil
}
