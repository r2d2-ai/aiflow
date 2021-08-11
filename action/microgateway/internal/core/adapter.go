package core

import (
	"github.com/r2d2-ai/aiflow/action/microgateway/api"
	"github.com/r2d2-ai/aiflow/activity"
)

// Adapter is an adapter activity for ServiceFunc
type Adapter struct {
	Handler api.ServiceFunc
}

// Metadata returns the metadata for the adapter activity
func (a *Adapter) Metadata() *activity.Metadata {
	return nil
}

// Eval evaluates the adapter activity
func (a *Adapter) Eval(ctx activity.Context) (done bool, err error) {
	return a.Handler(ctx)
}
