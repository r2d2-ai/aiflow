package log

import (
	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data/coerce"
	cv "gocv.io/x/gocv"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Input struct {
	Image *cv.Mat       `md:"image"` // The message to log
	ROIs  []interface{} `md:"rois"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image": i.Image,
		"rois":  i.ROIs,
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
	ctx.Logger().Infof("Image with size %s", input.Image.Size())
	// msg := input.Message

	// if input.AddDetails {
	// 	msg = fmt.Sprintf("'%s' - HostID [%s], HostName [%s], Activity [%s]", msg,
	// 		ctx.ActivityHost().ID(), ctx.ActivityHost().Name(), ctx.Name())
	// }

	// if input.UsePrint {
	// 	fmt.Println(msg)
	// } else {
	// 	ctx.Logger().Info(msg)
	// }

	return true, nil
}
