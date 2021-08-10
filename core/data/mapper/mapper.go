package mapper

import "github.com/r2d2-ai/aiflow/core/data"

type Factory interface {
	NewMapper(mappings map[string]interface{}) (Mapper, error)
}

// Mapper is a constructs that maps values to a map
type Mapper interface {
	Apply(scope data.Scope) (map[string]interface{}, error)
}
