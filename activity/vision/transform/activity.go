package transform

import (
	"image"
	"image/color"
	"math"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data/metadata"
	gocv "gocv.io/x/gocv"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{settings: s}
	return act, nil
}

// Activity is an Activity that is used to
// inputs : {image, rois}
// outputs: none
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)
	src := input.Image
	if src.Empty() {
		//TODO error handling
		return false, nil
	}
	dst := gocv.NewMat()
	s := a.settings
	inputWidth := src.Cols()
	inputHeight := src.Rows()
	outputWidth := s.OutputWidth
	outputHeight := s.OutputHeight

	if inputHeight != outputHeight || inputWidth != outputWidth {

		if s.ScaleMode == "Stretch" {
			var interp gocv.InterpolationFlags
			if inputWidth > outputWidth && inputHeight > outputHeight {
				interp = gocv.InterpolationArea
			} else {
				interp = gocv.InterpolationLinear
			}
			gocv.Resize(src, &dst, image.Pt(outputWidth, outputHeight), 0, 0, interp)
		} else {
			scale := math.Min(float64(outputWidth)/float64(inputWidth), float64(outputHeight)/float64(inputHeight))
			targetWidth := int(float64(inputWidth) * scale)
			targetHeight := int(float64(inputHeight) * scale)
			var interp gocv.InterpolationFlags
			if scale < 1 {
				interp = gocv.InterpolationArea
			} else {
				interp = gocv.InterpolationLinear
			}

			if s.ScaleMode == "Fit" {
				tmp := gocv.NewMat()
				defer tmp.Close()
				gocv.Resize(src, &tmp, image.Pt(targetWidth, targetHeight), 0, 0, interp)
				top := int((outputHeight - targetHeight) / 2)
				bottom := outputHeight - targetHeight - top
				left := int((outputWidth - targetWidth) / 2)
				right := outputWidth - targetWidth - left
				gocv.CopyMakeBorder(tmp, &dst, top, bottom, left, right, gocv.BorderConstant, color.RGBA{})
			} else {
				gocv.Resize(src, &dst, image.Pt(targetWidth, targetHeight), 0, 0, interp)
			}
		}
		src = dst
	}

	output := &Output{}
	output.Image = dst
	ctx.SetOutputObject(output)

	//ctx.Logger().Infof("Image from %s - %s", input.GroupdId, input.CameraId)

	return true, nil
}
