package object_detection

import (
	"image"

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
	// if input.Image == nil {
	// 	//no image nothing to do
	// 	return true, nil
	// }

	img := input.Image

	if img.Empty() {
		//TODO error handling
		return false, nil
	}

	ratio := 1.0 / 127.5
	mean := gocv.NewScalar(127.5, 127.5, 127.5, 0)
	swapRGB := true
	blob := gocv.BlobFromImage(img, ratio, image.Pt(300, 300), mean, swapRGB, false)
	blob.Close()
	//ctx.Logger().Infof("Image from %s - %s", input.GroupdId, input.CameraId)

	return true, nil
}
