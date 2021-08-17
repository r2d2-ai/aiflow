package object_detection

import (
	"github.com/r2d2-ai/aiflow/data/coerce"
	vision_types "github.com/r2d2-ai/aiflow/data/vision_types"
	gocv "gocv.io/x/gocv"
)

type Settings struct {
	Model string `md:"model,required"`
}

type Input struct {
	Image    *gocv.Mat           `md:"image"`
	ROIs     *[]vision_types.ROI `md:"rois"`
	GroupdId string              `md:"groupId"`
	CameraId string              `md:"cameraId"`
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

	i.ROIs, err = coerce.ToROIs(values["rois"])
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

type Output struct {
	Results map[string]interface{} `md:"results"` //
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"results": o.Results,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Results, err = coerce.ToObject(values["results"])
	if err != nil {
		return err
	}

	return nil
}
