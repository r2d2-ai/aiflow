package api

import (
	"github.com/r2d2-ai/aiflow/action/microgateway"
	microapi "github.com/r2d2-ai/aiflow/action/microgateway/api"
	"github.com/r2d2-ai/aiflow/activity/common/jsexec"
	"github.com/r2d2-ai/aiflow/api"
	"github.com/r2d2-ai/aiflow/engine"
	trigger "github.com/r2d2-ai/aiflow/trigger/net/rest"
)

// Example returns an API example
func Example() (engine.Engine, error) {
	app := api.NewApp()

	gateway := microapi.New("JS")

	service := gateway.NewService("JS", &jsexec.Activity{})
	service.SetDescription("Calculate sum")
	service.AddSetting("script", "result.sum = parameters.a + parameters.b")

	step := gateway.NewStep(service)
	step.AddInput("parameters", map[string]interface{}{"a": 1.0, "b": 2.0})

	response := gateway.NewResponse(false)
	response.SetCode(200)
	response.SetData("=$.JS.outputs.result")
	settings, err := gateway.AddResource(app)
	if err != nil {
		return nil, err
	}

	trg := app.NewTrigger(&trigger.Trigger{}, &trigger.Settings{Port: 9096})
	handler, err := trg.NewHandler(&trigger.HandlerSettings{
		Method: "GET",
		Path:   "/calculate",
	})
	if err != nil {
		return nil, err
	}

	_, err = handler.NewAction(&microgateway.Action{}, settings)
	if err != nil {
		return nil, err
	}

	return api.NewEngine(app)
}

func main() {
	e, err := Example()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
