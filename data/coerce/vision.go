package coerce

import (
	"fmt"
	"image"

	vision_types "github.com/r2d2-ai/aiflow/data/vision_types"
	gocv "gocv.io/x/gocv"
)

func ToImage(val interface{}) (gocv.Mat, error) {
	switch t := val.(type) {
	case *gocv.Mat:
		return *t, nil
	case gocv.Mat:
		return t, nil
	case nil:
		return gocv.Mat{}, nil
	default:
		// try to create config from map[string]interface{}
		return gocv.Mat{}, fmt.Errorf("unable to create image from '%#v'", val)
	}

}

func ToPoint(val interface{}) (image.Point, error) {
	switch t := val.(type) {
	case *image.Point:
		return *t, nil
	case image.Point:
		return t, nil
	case nil:
		return image.Point{}, nil
	default:
		return image.Point{}, fmt.Errorf("unable to create point from '%#v'", val)
	}
}

func ToPoints(val interface{}) ([]image.Point, error) {
	switch t := val.(type) {
	case []image.Point:
		return t, nil
	case *[]image.Point:
		return *t, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to create points from '%#v'", val)
	}
}

func ToROI(val interface{}) (vision_types.ROI, error) {
	switch t := val.(type) {
	case *vision_types.ROI:
		return *t, nil
	case vision_types.ROI:
		return t, nil
	case nil:
		return vision_types.ROI{}, nil
	default:
		return vision_types.ROI{}, fmt.Errorf("unable to create roi from '%#v'", val)
	}
}

func ToROIs(val interface{}) ([]vision_types.ROI, error) {
	switch t := val.(type) {
	case []vision_types.ROI:
		return t, nil
	case *[]vision_types.ROI:
		return *t, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to create rois from '%#v'", val)
	}
}
