package coerce

import (
	"fmt"

	cv "gocv.io/x/gocv"
)

func ToImage(val interface{}) (*cv.Mat, error) {
	switch t := val.(type) {
	case *cv.Mat:
		return t, nil

	default:
		// try to create config from map[string]interface{}
		return nil, fmt.Errorf("unable to create mat from '%#v'", val)
	}

}
