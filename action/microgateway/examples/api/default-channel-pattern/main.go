package main

import (
	_ "github.com/r2d2-ai/aiflow/action/microgateway/activity/circuitbreaker"
	_ "github.com/r2d2-ai/aiflow/action/microgateway/activity/jwt"
	"github.com/r2d2-ai/aiflow/action/microgateway/examples"
	"github.com/r2d2-ai/aiflow/engine"
	_ "github.com/r2d2-ai/contrib/activity/channel"
)

func main() {

	e, err := examples.DefaultChannelPattern()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
