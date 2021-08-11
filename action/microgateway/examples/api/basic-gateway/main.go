package main

import (
	"github.com/r2d2-ai/aiflow/action/microgateway/examples"
	"github.com/r2d2-ai/aiflow/engine"
)

func main() {
	e, err := examples.BasicGatewayExample()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
