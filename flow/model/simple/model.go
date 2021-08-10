package simple

import (
	"github.com/r2d2-ai/AIflow/flow/model"
)

const (
	ModelName = "AIflow-simple"
)

func init() {
	model.Register(New())
}

func New() *model.FlowModel {
	m := model.New(ModelName)
	m.RegisterFlowBehavior(&FlowBehavior{})
	m.RegisterDefaultTaskBehavior("basic", &TaskBehavior{})
	m.RegisterTaskBehavior("iterator", &IteratorTaskBehavior{})
	m.RegisterTaskBehavior("doWhile", &DoWhileTaskBehavior{})
	return m
}
