package pipeline

import (
	"encoding/json"
	"fmt"

	"github.com/r2d2-ai/aiflow/app/resource"
	"github.com/r2d2-ai/aiflow/data/mapper"
	"github.com/r2d2-ai/aiflow/data/resolve"
)

const (
	ResType    = "stream"
	ResTypeOld = "pipeline"
)

func NewResourceLoader(mapperFactory mapper.Factory, resolver resolve.CompositeResolver) resource.Loader {
	return &ResourceLoader{mapperFactory: mapperFactory, resolver: resolver}
}

type ResourceLoader struct {
	mapperFactory mapper.Factory
	resolver      resolve.CompositeResolver
}

func (rl *ResourceLoader) LoadResource(config *resource.Config) (*resource.Resource, error) {

	var pipelineCfgBytes []byte

	pipelineCfgBytes = config.Data

	var defConfig *DefinitionConfig
	err := json.Unmarshal(pipelineCfgBytes, &defConfig)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling stream resource with id '%s', %s", config.ID, err.Error())
	}

	defConfig.id = config.ID

	pDef, err := NewDefinition(defConfig, rl.mapperFactory, rl.resolver)
	if err != nil {
		return nil, err
	}

	return resource.New(ResType, pDef), nil
}
