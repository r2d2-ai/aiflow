package engine

import (
	"fmt"
	"strings"

	"github.com/r2d2-ai/aiflow/action"
	"github.com/r2d2-ai/aiflow/app"
	"github.com/r2d2-ai/aiflow/data/property"
	"github.com/r2d2-ai/aiflow/engine/channels"
	"github.com/r2d2-ai/aiflow/engine/runner"
	"github.com/r2d2-ai/aiflow/engine/secret"
	"github.com/r2d2-ai/aiflow/support/log"
	"github.com/r2d2-ai/aiflow/support/managed"
	"github.com/r2d2-ai/aiflow/support/service"
	"github.com/r2d2-ai/aiflow/support/trace"
)

// engineImpl is the type for the Default Engine Implementation
type engineImpl struct {
	config         *Config
	app            *app.App
	actionRunner   action.Runner
	serviceManager *service.Manager
	logger         log.Logger
}

type Option func(*engineImpl) error

// New creates a new Engine
func New(appConfig *app.Config, options ...Option) (Engine, error) {
	if appConfig == nil {
		return nil, fmt.Errorf("no App configuration provided")
	}
	if len(appConfig.Name) == 0 {
		return nil, fmt.Errorf("no App name provided")
	}
	if len(appConfig.Version) == 0 {
		return nil, fmt.Errorf("no App version provided")
	}

	engine := &engineImpl{}
	logger := log.ChildLogger(log.RootLogger(), "engine")
	engine.logger = logger

	// Register tracer as managed service
	if trace.Enabled() {
		LifeCycle(trace.GetTracer())
	}
	//log.SetLogLevel(log.DebugLevel, logger)

	for _, option := range options {
		err := option(engine)
		if err != nil {
			return nil, err
		}
	}

	if engine.config == nil {
		config := &Config{}
		config.StopEngineOnError = StopEngineOnError()
		config.RunnerType = GetRunnerType()
		engine.config = config
	}

	if engine.actionRunner == nil {
		var actionRunner action.Runner

		runnerType := engine.config.RunnerType
		if strings.EqualFold(ValueRunnerTypePooled, runnerType) {
			actionRunner = runner.NewPooled(NewPooledRunnerConfig())
		} else if strings.EqualFold(ValueRunnerTypeDirect, runnerType) {
			actionRunner = runner.NewDirect()
		} else {
			return nil, fmt.Errorf("unknown runner type: %s", runnerType)
		}

		logger.Debugf("Using '%s' Action Runner", runnerType)
		engine.actionRunner = actionRunner
	}

	var appOptions []app.Option
	if !engine.config.StopEngineOnError {
		appOptions = append(appOptions, app.ContinueOnError)
	}

	// Setup Property Resolvers
	propResolvers := GetAppPropertyValueResolvers(logger)
	enablePropertiesResolution := false
	if len(propResolvers) > 0 {
		err := property.EnableExternalPropertyResolvers(propResolvers)
		if err != nil {
			return nil, err
		}

		enablePropertiesResolution = true
	}

	// properties post processors (properties resolver if enabled, secret properties replacer)
	var postProcessors []property.PostProcessor
	if enablePropertiesResolution {
		postProcessors = append(postProcessors, property.ExternalResolverProcessor)
	}
	postProcessors = append(postProcessors, secret.PropertyProcessor)

	appOptions = append(appOptions, app.FinalizeProperties(postProcessors...))

	engine.serviceManager = service.NewServiceManager()
	appOptions = append(appOptions, app.EngineSettings(engine.serviceManager, engine.config.ActionSettings))

	//setup services
	if len(engine.config.Services) > 0 {
		for _, sConfig := range engine.config.Services {
			if sConfig.Enabled {
				f := service.GetFactory(sConfig.Ref)
				if f != nil {
					svc, err := f.NewService(&service.Config{
						Settings: sConfig.Settings,
					})
					if err != nil {
						return nil, err
					}
					err = engine.serviceManager.RegisterService(svc)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	// Create the application
	flowApp, err := app.New(appConfig, engine.actionRunner, appOptions...)
	if err != nil {
		return nil, err
	}

	logger.Debugf("Creating app [ %s ] with version [ %s ]", appConfig.Name, appConfig.Version)
	engine.app = flowApp

	return engine, nil
}

func ConfigOption(engineJson string, compressed bool) func(*engineImpl) error {
	return func(e *engineImpl) error {

		cfg, err := LoadEngineConfig(engineJson, compressed)
		if err != nil {
			return err
		}

		e.config = cfg

		return nil
	}
}

func (e *engineImpl) App() *app.App {
	return e.app
}

//Start initializes and starts the Triggers and initializes the Actions
func (e *engineImpl) Start() error {

	logger := e.logger

	logger.Infof("Starting app [ %s ] with version [ %s ]", e.app.Name(), e.app.Version())

	logger.Info("Engine Starting...")

	logger.Info("Starting Services...")

	actionRunner := e.actionRunner.(interface{})

	if managedRunner, ok := actionRunner.(managed.Managed); ok {
		_ = managed.Start("ActionRunner Service", managedRunner)
	}

	err := e.serviceManager.Start()

	if err != nil {
		logger.Error("Error Starting Services - " + err.Error())
	} else {
		logger.Info("Started Services")
	}

	if len(managedServices) > 0 {
		for _, mService := range managedServices {
			err = mService.Start()
			if err != nil {
				logger.Error("Error Starting Services - " + err.Error())
				//TODO Should we exit here?
			}
		}
	}

	logger.Info("Starting Application...")
	e.app.PostAppEvent(app.STARTING)
	err = e.app.Start()
	if err != nil {
		e.app.PostAppEvent(app.FAILED)
		return err
	}
	e.app.PostAppEvent(app.STARTED)
	logger.Info("Application Started")

	if channels.Count() > 0 {
		logger.Info("Starting Engine Channels...")
		_ = channels.Start()
		logger.Info("Engine Channels Started")
	}

	logger.Info("Engine Started")

	return nil
}

func (e *engineImpl) Stop() error {

	logger := e.logger

	logger.Info("Engine Stopping...")

	if channels.Count() > 0 {
		logger.Info("Stopping Engine Channels...")
		_ = channels.Stop()
		logger.Info("Engine Channels Stopped...")
	}

	logger.Info("Stopping Application...")
	e.app.PostAppEvent(app.STOPPING)
	_ = e.app.Stop()
	logger.Info("Application Stopped")
	e.app.PostAppEvent(app.STOPPED)

	//TODO temporarily add services
	logger.Info("Stopping Services...")

	actionRunner := e.actionRunner.(interface{})

	if managedRunner, ok := actionRunner.(managed.Managed); ok {
		_ = managed.Stop("ActionRunner", managedRunner)
	}

	err := e.serviceManager.Stop()

	if err != nil {
		logger.Error("Error Stopping Services - " + err.Error())
	} else {
		logger.Info("Stopped Services")
	}

	if len(managedServices) > 0 {
		for _, mService := range managedServices {
			err = mService.Stop()
			if err != nil {
				logger.Error("Error Stopping Services - " + err.Error())
			}
		}
	}

	logger.Info("Engine Stopped")
	log.Sync()

	return nil
}
