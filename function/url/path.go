package url

import (
	"fmt"
	"net/url"

	"github.com/r2d2-ai/aiflow/data"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/data/expression/function"
)

func init() {
	_ = function.Register(&fnPath{})
}

type fnPath struct {
}

// Name returns the name of the function
func (fnPath) Name() string {
	return "path"
}

// Sig returns the function signature
func (fnPath) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval executes the function
func (fnPath) Eval(params ...interface{}) (interface{}, error) {
	rawString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to string: %s", params[0], err.Error())
	}
	u, err := url.Parse(rawString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse url: %s", err.Error())
	}
	return u.Path, nil
}
