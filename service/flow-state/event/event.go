package event

import (
	"time"

	"github.com/r2d2-ai/aiflow/action/flow/state"
	coreevent "github.com/r2d2-ai/aiflow/engine/event"
)

const EventType = "streamingStepEvent"

type stepEvent struct {
	time time.Time
	step *state.Step
}

// Step retuns the step date
func (re *stepEvent) Step() *state.Step {
	return re.step
}

// Returns event time
func (re *stepEvent) Time() time.Time {
	return re.time
}

func PostStepEvent(step *state.Step) {
	if coreevent.HasListener(EventType) {
		fe := &stepEvent{
			time: time.Now(),
			step: step,
		}
		coreevent.Post(EventType, fe)
	}
}
