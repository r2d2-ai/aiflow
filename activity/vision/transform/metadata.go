package transform

import (
	"github.com/r2d2-ai/aiflow/data/coerce"
	gocv "gocv.io/x/gocv"
)

type Settings struct {
	OutputWidth      int    `md:"output_width"`
	OutputHeight     int    `md:"output_height"`
	Rotation         int    `md:"rotation"`
	FlipHorizontally bool   `md:"flip_horizontally"`
	FlipVertically   bool   `md:"flip_vertically"`
	ScaleMode        string `md:"scale_mode,allowed(Stretch,Fit,Fill,Crop)"`
}

type Input struct {
	Image gocv.Mat `md:"image"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image": i.Image,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Image, err = coerce.ToImage(values["image"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	Image            gocv.Mat      `md:"image"`
	LetterboxPadding []interface{} `md:"letterbox_padding"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image":             o.Image,
		"letterbox_padding": o.LetterboxPadding,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Image, err = coerce.ToImage(values["image"])
	if err != nil {
		return err
	}

	o.LetterboxPadding, err = coerce.ToArray(values["letterbox_padding"])
	if err != nil {
		return err
	}
	return nil
}
