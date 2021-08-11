package main

import (
	"github.com/r2d2-ai/aiflow/action/microgateway/activity/ratelimiter/examples"
	"github.com/r2d2-ai/aiflow/engine"
)

func main() {
	e, err := examples.Example("3-M", 0)
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
