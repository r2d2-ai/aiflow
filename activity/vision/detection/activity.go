package detection

import (
	"image"

	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data/coerce"
	cv "gocv.io/x/gocv"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Input struct {
	Image    *cv.Mat       `md:"image"` // The message to log
	ROIs     []interface{} `md:"rois"`
	GroupdId string        `md:"groupId"`
	CameraId string        `md:"cameraId"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image":    i.Image,
		"rois":     i.ROIs,
		"groupId":  i.GroupdId,
		"cameraId": i.CameraId,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Image, err = coerce.ToImage(values["image"])
	if err != nil {
		return err
	}

	i.ROIs, err = coerce.ToArray(values["rois"])
	if err != nil {
		return err
	}

	i.GroupdId, err = coerce.ToString(values["groupId"])
	if err != nil {
		return err
	}

	i.CameraId, err = coerce.ToString(values["cameraId"])
	if err != nil {
		return err
	}

	return nil
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : {image, rois}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)
	if input.Image == nil {
		//no image nothing to do
		return true, nil
	}

	img := *input.Image
	ratio := 1.0 / 127.5
	mean := cv.NewScalar(127.5, 127.5, 127.5, 0)
	swapRGB := true
	blob := cv.BlobFromImage(img, ratio, image.Pt(300, 300), mean, swapRGB, false)
	blob.Close()
	ctx.Logger().Infof("Image from %s - %s", input.GroupdId, input.CameraId)

	return true, nil
}
