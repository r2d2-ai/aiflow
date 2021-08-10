package ondemand

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/r2d2-ai/aiflow/core/action"
	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/data/coerce"
	"github.com/r2d2-ai/aiflow/core/data/expression"
	"github.com/r2d2-ai/aiflow/core/data/mapper"
	"github.com/r2d2-ai/aiflow/core/data/metadata"
	"github.com/r2d2-ai/aiflow/core/data/resolve"
	"github.com/r2d2-ai/aiflow/core/support"
	"github.com/r2d2-ai/aiflow/core/support/log"
	"github.com/r2d2-ai/aiflow/core/support/service"
	"github.com/r2d2-ai/aiflow/flow/definition"
	"github.com/r2d2-ai/aiflow/flow/instance"
	"github.com/r2d2-ai/aiflow/flow/model"
	"github.com/r2d2-ai/aiflow/flow/model/simple"
	"github.com/r2d2-ai/aiflow/flow/state"
	flowsupport "github.com/r2d2-ai/aiflow/flow/support"
)

const (
	ivFlowPackage = "flowPackage"
)

const (
	StateRecordingMode = "stateRecordingMode"
	// Deprecated
	RtSettingStepMode     = "stepRecordingMode"
	RtSettingSnapshotMode = "snapshotRecordingMode"
)

type FlowAction struct {
	FlowURI    string
	IoMetadata *metadata.IOMetadata
}

type FlowPackage struct {
	Inputs  map[string]interface{}    `json:"inputs"`
	Outputs map[string]interface{}    `json:"outputs"`
	Flow    *definition.DefinitionRep `json:"flow"`
}

var idGenerator *support.Generator
var record bool
var flowManager *flowsupport.FlowManager
var logger log.Logger
var stateRecorder state.Recorder
var stateRecordingMode = state.RecordingModeOff

var actionMd = action.ToMetadata(&Settings{})

//todo expose and support this properly
var maxStepCount = 1000000

type Settings struct {
}

func init() {
	_ = action.Register(&FlowAction{}, &ActionFactory{})
}

type ActionFactory struct {
}

func (f *ActionFactory) Initialize(ctx action.InitContext) error {

	logger = log.ChildLogger(log.RootLogger(), "flow")

	if flowManager != nil {
		return nil
	}

	sm := ctx.ServiceManager()

	srService := sm.FindService(func(s service.Service) bool {
		_, ok := s.(state.Recorder)
		return ok
	})

	if len(ctx.RuntimeSettings()) > 0 {
		mode, ok := ctx.RuntimeSettings()[StateRecordingMode]
		if !ok {
			// For backward compatible
			sStepMode := ctx.RuntimeSettings()[RtSettingStepMode]
			sSnapshotMode := ctx.RuntimeSettings()[RtSettingSnapshotMode]

			stepMode, _ := coerce.ToString(sStepMode)
			snapshotMode, _ := coerce.ToString(sSnapshotMode)

			recordSteps := strings.EqualFold("full", stepMode)
			recordSnapshot := strings.EqualFold("full", snapshotMode)
			if recordSteps && recordSnapshot {
				stateRecordingMode = state.RecordingModeFull
			} else if recordSteps {
				stateRecordingMode = state.RecordingModeStep
			} else if recordSnapshot {
				stateRecordingMode = state.RecordingModeSnapshot
			} else {
				stateRecordingMode = state.RecordingModeOff
			}
		} else {
			var err error
			stateRecordingMode, err = state.ToRecordingMode(mode)
			if err != nil {
				return nil
			}
		}
	}

	if srService != nil {
		stateRecorder = srService.(state.Recorder)
		if state.RecordSteps(stateRecordingMode) {
			instance.EnableChangeTracking(true, stateRecordingMode)
		}
	}

	exprFactory := expression.NewFactory(definition.GetDataResolver())
	mapperFactory := mapper.NewFactory(definition.GetDataResolver())

	definition.SetMapperFactory(mapperFactory)
	definition.SetExprFactory(exprFactory)

	if idGenerator == nil {
		idGenerator, _ = support.NewGenerator()
	}

	//todo fix the following
	model.RegisterDefault(simple.New())
	flowManager = flowsupport.NewFlowManager(nil)
	flowsupport.InitDefaultDefLookup(flowManager, ctx.ResourceManager())

	return nil

}

//Metadata get the Action's metadata
func (fa *FlowAction) Metadata() *action.Metadata {
	return actionMd
}

func (fa *FlowAction) IOMetadata() *metadata.IOMetadata {
	return fa.IoMetadata
}

func (f *ActionFactory) New(config *action.Config) (action.Action, error) {

	flowAction := &FlowAction{}

	if config.Settings != nil {

	} else {
		flowAction.IoMetadata = &metadata.IOMetadata{Input: nil, Output: nil}
	}

	//temporary hack to support dynamic process running by tester
	//if config.Data == nil {
	//	return flowAction, nil
	//}

	return flowAction, nil

}

// Run implements action.Action.Run
func (fa *FlowAction) Run(ctx context.Context, inputs map[string]interface{}, handler action.ResultHandler) error {

	logger.Info("Running OnDemand Flow Action")

	fpAttr, exists := inputs[ivFlowPackage]

	flowPackage := &FlowPackage{}

	if exists {

		raw := fpAttr.(json.RawMessage)
		err := json.Unmarshal(raw, flowPackage)
		if err != nil {
			return err
		}
	} else {
		return errors.New("flow package not provided")
	}

	logger.Debugf("InputMappings: %+v", flowPackage.Inputs)
	logger.Debugf("OutputMappings: %+v", flowPackage.Outputs)

	flowDef, err := definition.NewDefinition(flowPackage.Flow)

	if err != nil {
		return err
	}

	flowInputs, err := ApplyMappings(flowPackage.Inputs, inputs)
	if err != nil {
		return err
	}

	instanceID := idGenerator.NextAsString()
	logger.Debug("Creating Flow Instance: ", instanceID)

	inst, err := instance.NewIndependentInstance(instanceID, "", flowDef, logger)

	if err != nil {
		return err
	}

	logger.Debugf("Executing Flow Instance: %s", inst.ID())

	inst.Start(flowInputs)

	stepCount := 0
	hasWork := true

	inst.SetResultHandler(handler)

	go func() {

		defer handler.Done()

		if !inst.FlowDefinition().ExplicitReply() {

			results := make(map[string]interface{})

			results["id"] = inst.ID

			handler.HandleResult(results, nil)
		}

		for hasWork && inst.Status() < model.FlowStatusCompleted && stepCount < maxStepCount {
			stepCount++
			logger.Debugf("Step: %d", stepCount)
			hasWork = inst.DoStep()

			if record {
				//ep.GetStateRecorder().RecordSnapshot(inst)
				//ep.GetStateRecorder().RecordStep(inst)
			}
		}

		if inst.Status() == model.FlowStatusCompleted {
			returnData, err := inst.GetReturnData()

			handler.HandleResult(returnData, err)

		} else if inst.Status() == model.FlowStatusFailed {
			handler.HandleResult(nil, inst.GetError())
		}

		logger.Debugf("Done Executing flow instance [%s] - Status: %d", inst.ID(), inst.Status())

		if inst.Status() == model.FlowStatusCompleted {
			logger.Infof("Flow instance [%s] Completed Successfully", inst.ID())
		} else if inst.Status() == model.FlowStatusFailed {
			logger.Infof("Flow instance [%s] Failed", inst.ID())
		}
	}()

	return nil
}

func ApplyMappings(mappings map[string]interface{}, inputs map[string]interface{}) (map[string]interface{}, error) {

	mapperFactory := mapper.NewFactory(resolve.GetBasicResolver())

	m, err := mapperFactory.NewMapper(mappings)
	if err != nil {
		return nil, err
	}

	inScope := data.NewSimpleScope(inputs, nil)

	out, err := m.Apply(inScope)

	return out, nil
}
