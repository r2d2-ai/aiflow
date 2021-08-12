package postgres

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/r2d2-ai/aiflow/action/flow/state"
	"github.com/r2d2-ai/aiflow/data/coerce"
	"github.com/r2d2-ai/aiflow/service/flow-state/store/metadata"
)

func NewStore(settings map[string]interface{}) (*StepStore, error) {
	db, err := NewDB(settings)
	if err != nil {
		return nil, err
	}
	return &StepStore{db: &StatefulDB{db: db}}, err
}

type StepStore struct {
	sync.RWMutex
	db             *StatefulDB
	userId         string
	appId          string
	stepContainers map[string]*stepContainer
	snapshots      sync.Map
}

func (s *StepStore) GetStatus(flowId string) int {
	s.RLock()
	sc, ok := s.stepContainers[flowId]
	s.RUnlock()

	if ok {
		return sc.Status()
	}
	return -1
}

//func (s *StepStore) GetFlow(flowId string) *state.FlowInfo {
//
//	s.RLock()
//	sc, ok := s.stepContainers[flowId]
//	s.RUnlock()
//
//	if ok {
//		return &state.FlowInfo{Id: flowId, Status: sc.Status(), FlowURI: sc.flowURI}
//	}
//
//	return nil
//}

func (s *StepStore) GetFailedFlows(metadata *metadata.Metadata) ([]*state.FlowInfo, error) {
	var whereStr = "where"
	if len(metadata.Username) < 0 {
		return nil, fmt.Errorf("unauthorized, please provide user infomation")
	}
	whereStr += " userId='" + metadata.Username + "'"

	if len(metadata.AppId) < 0 {
		return nil, fmt.Errorf("please provide flow id or flow name")
	}
	whereStr += "  and appId='" + metadata.AppId + "'"

	if len(metadata.HostId) > 0 {
		whereStr += "  and hostId='" + metadata.HostId + "'"
	}

	if len(metadata.FlowName) > 0 {
		whereStr += "  and flowname='" + metadata.FlowName + "'"
	}

	whereStr += " and status = 'Failed'"

	set, err := s.db.query("select flowinstanceid, flowname, status from flowstate "+whereStr, nil)
	if err != nil {
		return nil, err
	}

	v, _ := json.Marshal(set.Record)
	fmt.Println(string(v))

	var flowinfo []*state.FlowInfo
	for _, v := range set.Record {
		m := *v
		id, _ := coerce.ToString(m["flowinstanceid"])
		flowName, _ := coerce.ToString(m["flowname"])
		status, _ := coerce.ToString(m["status"])
		info := &state.FlowInfo{
			Id:         id,
			FlowName:   flowName,
			FlowStatus: status,
		}
		flowinfo = append(flowinfo, info)
	}
	return flowinfo, err
}

func (s *StepStore) GetFlows(metadata *metadata.Metadata) ([]*state.FlowInfo, error) {
	var whereStr = "where"
	if len(metadata.Username) < 0 {
		return nil, fmt.Errorf("unauthorized, please provide user infomation")
	}
	whereStr += " userId='" + metadata.Username + "'"

	if len(metadata.AppId) < 0 {
		return nil, fmt.Errorf("please provide flow id or flow name")
	}
	whereStr += "  and appId='" + metadata.AppId + "'"

	if len(metadata.HostId) > 0 {
		whereStr += "  and hostId='" + metadata.HostId + "'"
	}

	if len(metadata.FlowName) > 0 {
		whereStr += "  and flowname='" + metadata.FlowName + "'"
	}

	set, err := s.db.query("select flowinstanceid, flowname, status from flowstate "+whereStr, nil)
	if err != nil {
		return nil, err
	}

	v, _ := json.Marshal(set.Record)
	fmt.Println(string(v))

	var flowinfo []*state.FlowInfo
	for _, v := range set.Record {
		m := *v
		id, _ := coerce.ToString(m["flowinstanceid"])
		flowName, _ := coerce.ToString(m["flowname"])
		status, _ := coerce.ToString(m["status"])
		info := &state.FlowInfo{
			Id:         id,
			FlowName:   flowName,
			FlowStatus: status,
		}
		flowinfo = append(flowinfo, info)
	}
	return flowinfo, err
}

func (s *StepStore) GetFlow(flowid string, metadata *metadata.Metadata) (*state.FlowInfo, error) {
	var whereStr = "where flowinstanceid = '" + flowid + "'"
	if len(metadata.Username) > 0 {
		whereStr += " and userId='" + metadata.Username + "'"
	}

	if len(metadata.AppId) > 0 {
		whereStr += "  and appId='" + metadata.AppId + "'"
	}

	if len(metadata.HostId) > 0 {
		whereStr += "  and hostId='" + metadata.HostId + "'"
	}

	set, err := s.db.query("select flowinstanceid, flowname, status from flowstate "+whereStr, nil)
	if err != nil {
		return nil, err
	}

	v, _ := json.Marshal(set.Record)
	fmt.Println(string(v))

	var flowinfo []*state.FlowInfo
	for _, v := range set.Record {
		m := *v
		id, _ := coerce.ToString(m["flowinstanceid"])
		flowName, _ := coerce.ToString(m["flowname"])
		status, _ := coerce.ToString(m["status"])
		info := &state.FlowInfo{
			Id:         id,
			FlowName:   flowName,
			FlowStatus: status,
			FlowURI:    "res://flow:" + flowName,
		}
		flowinfo = append(flowinfo, info)
	}
	if len(flowinfo) <= 0 {
		return nil, fmt.Errorf("flow details [%s] not found", flowid)
	}
	return flowinfo[0], err
}

func (s *StepStore) SaveStep(step *state.Step) error {
	_, err := s.db.InsertSteps(step)
	return err
}

func (s *StepStore) GetSteps(flowId string) ([]*state.Step, error) {

	set, err := s.db.query("select stepdata from steps where flowinstanceid = '"+flowId+"'", nil)
	if err != nil {
		return nil, err
	}

	var steps []*state.Step
	for _, v := range set.Record {
		m := *v
		s1, err := coerce.ToBytes(m["stepdata"])
		if err != nil {
			return nil, fmt.Errorf("decodeBase64 for step data error:", err.Error())
		}
		dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(s1)))
		n, err := base64.StdEncoding.Decode(dbuf, s1)
		stePdata := dbuf[:n]
		var step *state.Step
		err = json.Unmarshal(stePdata, &step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}

	if len(steps) <= 0 {
		return nil, fmt.Errorf("step for flow instance [%s] not found", flowId)
	}
	return steps, err
}

func (s *StepStore) GetStepsNoData(flowId string) ([]map[string]string, error) {

	set, err := s.db.query("select stepid, taskname, status from steps where flowinstanceid = '"+flowId+"'", nil)
	if err != nil {
		return nil, err
	}

	var steps []map[string]string
	for _, v := range set.Record {
		m := *v
		stepData := make(map[string]string)
		s1, _ := coerce.ToString(m["stepid"])
		stepData["stepId"] = s1

		status, _ := coerce.ToString(m["status"])
		stepData["status"] = status

		name, _ := coerce.ToString(m["taskname"])
		stepData["taskName"] = name
		steps = append(steps, stepData)
	}

	if len(steps) <= 0 {
		return nil, fmt.Errorf("step for flow instance [%s] not found", flowId)
	}
	return steps, err
}

func (s *StepStore) Delete(flowId string) {
	s.Lock()
	delete(s.stepContainers, flowId)
	s.Unlock()
}

type stepContainer struct {
	sync.RWMutex
	status  int
	flowURI string
	steps   []*state.Step
}

func (sc *stepContainer) Status() int {
	sc.RLock()
	status := sc.status
	sc.RUnlock()

	return status
}

func (sc *stepContainer) AddStep(step *state.Step) {
	sc.Lock()

	if len(step.FlowChanges) > 0 {
		if step.FlowChanges[0] != nil && step.FlowChanges[0].SubflowId == 0 {
			if status := step.FlowChanges[0].Status; status != -1 {
				sc.status = status
			}
			if uri := step.FlowChanges[0].FlowURI; uri != "" {
				sc.flowURI = uri
			}
		}
	}

	sc.steps = append(sc.steps, step)
	sc.Unlock()
}

func (sc *stepContainer) Steps() []*state.Step {
	sc.RLock()
	steps := sc.steps
	sc.RUnlock()
	return steps
}

func (s *StepStore) SaveSnapshot(snapshot *state.Snapshot) error {
	//replaces existing snapshot
	return nil
}

func (s *StepStore) GetSnapshot(flowId string) *state.Snapshot {
	if snapshot, ok := s.snapshots.Load(flowId); ok {
		return snapshot.(*state.Snapshot)
	}
	return nil
}

func (s *StepStore) RecordStart(flowState *state.FlowState) error {
	_, err := s.db.InsertFlowState(flowState)
	return err
}

func (s *StepStore) RecordEnd(flowState *state.FlowState) error {
	_, err := s.db.UpdateFlowState(flowState)
	return err
}
