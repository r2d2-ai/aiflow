package main

import (
	"github.com/r2d2-ai/aiflow/action/microgateway/activity/anomaly/examples"
	"github.com/r2d2-ai/aiflow/engine"
)

func main() {
	e, err := examples.Example()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
