package test

import (
	"github.com/r2d2-ai/ai-box/flow/model"
	"github.com/r2d2-ai/ai-box/flow/model/simple"
)

func init() {
	model.Register(NewTestModel())
}

func NewTestModel() *model.FlowModel {
	m := model.New("test")
	m.RegisterFlowBehavior(&simple.FlowBehavior{})
	m.RegisterDefaultTaskBehavior("basic", &simple.TaskBehavior{})

	return m
}
