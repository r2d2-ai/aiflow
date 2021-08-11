package state

import "github.com/r2d2-ai/aiflow/action/flow/state/change"

type Step struct {
	Id           int                   `json:"id"`
	FlowId       string                `json:"flowId"`
	FlowChanges  map[int]*change.Flow  `json:"flowChanges"`
	QueueChanges map[int]*change.Queue `json:"queueChanges,omitempty"`
}
