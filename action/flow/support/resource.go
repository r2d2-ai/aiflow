package support

import (
	"encoding/json"
	"fmt"

	"github.com/r2d2-ai/aiflow/action/flow/definition"
	"github.com/r2d2-ai/aiflow/app/resource"
)

const (
	ResTypeFlow = "flow"
)

type FlowLoader struct {
}

func (*FlowLoader) LoadResource(config *resource.Config) (*resource.Resource, error) {
	var flowDefBytes []byte

	flowDefBytes = config.Data

	var defRep *definition.DefinitionRep
	err := json.Unmarshal(flowDefBytes, &defRep)
	if err != nil {
		return nil, fmt.Errorf("error loading flow resource with id '%s': %s", config.ID, err.Error())
	}

	flow, err := materializeFlow(defRep)
	if err != nil {
		return nil, err
	}

	return resource.New(ResTypeFlow, flow), nil
}
