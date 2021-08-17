package coerce

import (
	"fmt"

	vision_types "github.com/r2d2-ai/aiflow/data/vision_types"
	gocv "gocv.io/x/gocv"
)

func ToImage(val interface{}) (*gocv.Mat, error) {
	switch t := val.(type) {
	case *gocv.Mat:
		return t, nil
	case gocv.Mat:
		return &t, nil
	case nil:
		return nil, nil
	default:
		// try to create config from map[string]interface{}
		return nil, fmt.Errorf("unable to create image from '%#v'", val)
	}

}

func ToPoint(val interface{}) (gocv.Point2f, error) {
	switch t := val.(type) {
	case *gocv.Point2f:
		return *t, nil
	case gocv.Point2f:
		return t, nil
	default:
		return gocv.Point2f{}, fmt.Errorf("unable to create point from '%#v'", val)
	}
}

func ToPoints(val interface{}) (gocv.Point2fVector, error) {
	switch t := val.(type) {
	case []gocv.Point2f:
		pv := gocv.NewPoint2fVectorFromPoints(t)
		return pv, nil
	case *[]gocv.Point2f:
		pv := gocv.NewPoint2fVectorFromPoints(*t)
		return pv, nil
	case *gocv.Point2fVector:
		return *t, nil
	case gocv.Point2fVector:
		return t, nil
	default:
		return gocv.Point2fVector{}, fmt.Errorf("unable to create points from '%#v'", val)
	}
}

func ToROI(val interface{}) (*vision_types.ROI, error) {
	switch t := val.(type) {
	case *vision_types.ROI:
		return t, nil
	case vision_types.ROI:
		return &t, nil
	default:
		return nil, fmt.Errorf("unable to create roi from '%#v'", val)
	}
}

func ToROIs(val interface{}) (*[]vision_types.ROI, error) {
	switch t := val.(type) {
	case []vision_types.ROI:
		return &t, nil
	case *[]vision_types.ROI:
		return t, nil
	default:
		return nil, fmt.Errorf("unable to create rois from '%#v'", val)
	}
}
