package vision_types

import (
	"image"
)

type ROI struct {
	Name   string        `json:"name,omitempty"`
	Points []image.Point `json:"points"`
	Type   string        `json:"type"`
}
