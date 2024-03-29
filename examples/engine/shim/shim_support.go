package main

import (
	"os"

	_ "github.com/r2d2-ai/aiflow/data/expression/script"
	"github.com/r2d2-ai/aiflow/engine"
	"github.com/r2d2-ai/aiflow/support/log"
)

var (
	cfgJson       string
	cfgEngine     string
	cfgCompressed bool
	flowEngine    engine.Engine
)

func init() {
	log.SetLogLevel(log.RootLogger(), log.ErrorLevel)

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}

	flowEngine, err = engine.New(cfg, engine.ConfigOption(cfgEngine, cfgCompressed), engine.DirectRunner)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}
}
