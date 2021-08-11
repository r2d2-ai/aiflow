package streamtester

import (
	"testing"

	"github.com/r2d2-ai/aiflow/core/support"
	"github.com/r2d2-ai/aiflow/core/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {
	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}
