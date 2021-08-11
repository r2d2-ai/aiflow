package instance

import (
	"testing"

	"github.com/r2d2-ai/aiflow/action/flow/definition"
	"github.com/r2d2-ai/aiflow/core/data"
	"github.com/r2d2-ai/aiflow/core/data/resolve"
	"github.com/stretchr/testify/assert"
)

func TestWorkingDataScope_GetValue(t *testing.T) {

	vals := map[string]interface{}{"foo": 1, "bar": 2}
	baseScope := data.NewSimpleScope(vals, nil)

	workingDataScope := NewWorkingDataScope(baseScope)

	iteration := map[string]interface{}{
		"key":   1,
		"value": "blah",
	}

	workingDataScope.SetWorkingValue("iteration", iteration)

	v, exists := workingDataScope.GetWorkingValue("iteration")
	assert.True(t, exists)
	assert.Equal(t, iteration, v)

	v, exists = workingDataScope.GetValue("_W.iteration")
	assert.True(t, exists)
	assert.Equal(t, iteration, v)

	v, exists = workingDataScope.GetValue("foo")
	assert.True(t, exists)
	assert.Equal(t, 1, v)
}

func TestIterationResolver(t *testing.T) {

	vals := map[string]interface{}{"foo": 1, "bar": 2}
	baseScope := data.NewSimpleScope(vals, nil)

	workingDataScope := NewWorkingDataScope(baseScope)

	iteration := map[string]interface{}{
		"key":   1,
		"value": "blah",
	}

	workingDataScope.SetWorkingValue("iteration", iteration)

	resolver := resolve.NewCompositeResolver(map[string]resolve.Resolver{
		"iteration": &definition.IteratorResolver{}})

	v, err := resolver.Resolve("$iteration[key]", workingDataScope)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	v, err = resolver.Resolve("$iteration[value]", workingDataScope)
	assert.Nil(t, err)
	assert.Equal(t, "blah", v)
}
