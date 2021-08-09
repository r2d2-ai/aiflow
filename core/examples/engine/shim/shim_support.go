package main

import (
	"os"

	_ "github.com/r2d2-ai/ai-box/core/data/expression/script"
	"github.com/r2d2-ai/ai-box/core/engine"
	"github.com/r2d2-ai/ai-box/core/support/log"
)

var (
	cfgJson       string
	cfgEngine     string
	cfgCompressed bool
	AIflowEngine  engine.Engine
)

func init() {
	log.SetLogLevel(log.RootLogger(), log.ErrorLevel)

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}

	AIflowEngine, err = engine.New(cfg, engine.ConfigOption(cfgEngine, cfgCompressed), engine.DirectRunner)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}
}
