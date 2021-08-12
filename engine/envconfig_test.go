package engine

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestNewPooledConfigOk
func TestNewPooledConfigDefault(t *testing.T) {
	pooledConfig := NewPooledRunnerConfig()

	// assert Success
	assert.Equal(t, DefaultPooledWorkers, pooledConfig.NumWorkers)
	assert.Equal(t, DefaultPooledQueueSize, pooledConfig.WorkQueueSize)
}

//TestNewPooledConfigOk
func TestNewPooledConfigOverride(t *testing.T) {
	previousWorkers := os.Getenv(EnvKeyPooledWorkers)
	defer os.Setenv(EnvKeyPooledWorkers, previousWorkers)
	previousQueue := os.Getenv(EnvKeyPooledQueueSize)
	defer os.Setenv(EnvKeyPooledQueueSize, previousQueue)

	newWorkersValue := 6
	newQueueValue := 60

	// Change values
	_ = os.Setenv(EnvKeyPooledWorkers, strconv.Itoa(newWorkersValue))
	_ = os.Setenv(EnvKeyPooledQueueSize, strconv.Itoa(newQueueValue))

	pooledConfig := NewPooledRunnerConfig()

	// assert Success
	assert.Equal(t, newWorkersValue, pooledConfig.NumWorkers)
	assert.Equal(t, newQueueValue, pooledConfig.WorkQueueSize)
}
