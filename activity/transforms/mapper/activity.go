package activity_mapper

import (
	"github.com/r2d2-ai/aiflow/activity"
	"github.com/r2d2-ai/aiflow/data/mapper"
	"github.com/r2d2-ai/aiflow/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

type Settings struct {
	Mappings map[string]interface{} `md:"mappings,required"` // Set of mappings to execute
}

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{}

	ctx.Logger().Debugf("Mappings: %+v", s.Mappings)

	act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Activity is an Activity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type Activity struct {
	mapper mapper.Mapper
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	actionCtx := ctx.ActivityHost()

	if a.mapper == nil {
		//No mapping
		return true, nil
	}

	inputScope := actionCtx.Scope() //host data

	results, err := a.mapper.Apply(inputScope)
	if err != nil {
		return false, activity.NewError(err.Error(), "", nil)
	}

	for name, value := range results {
		err = actionCtx.Scope().SetValue(name, value)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
