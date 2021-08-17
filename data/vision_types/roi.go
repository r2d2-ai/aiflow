package vision_types

import (
	gocv "gocv.io/x/gocv"
)

type ROI struct {
	Name   string          `json:"name,omitempty"`
	Points *[]gocv.Point2f `json:"points"`
	Type   string          `json:"type"`
}
