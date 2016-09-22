package test

import (
	"strings"
	"testing"

	"github.com/foolbread/superlistener/protocol"
)

func Test_ProcessStateStarting(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:STOPPED tries:0"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_STARTING, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateStarting)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "STOPPED") {
		t.Error("from_state parse error!")
	}

	if pa.Tries != 0 {
		t.Error("tries parse error!")
	}
}

func Test_ProcessStateRunning(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:STARTING pid:2766"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_RUNNING, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateRunning)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "STARTING") {
		t.Error("from_state parse error!")
	}

	if pa.Pid != 2766 {
		t.Error("pid parse error!")
	}
}

func Test_ProcessStateStopping(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:STARTING pid:2766"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_STOPPING, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateStopping)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "STARTING") {
		t.Error("from_state parse error!")
	}

	if pa.Pid != 2766 {
		t.Error("pid parse error!")
	}
}

func Test_ProcessStateExited(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:RUNNING expected:0 pid:2766"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_EXITED, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateExited)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "RUNNING") {
		t.Error("from_state parse error!")
	}

	if pa.Expected != 0 {
		t.Error("expected parse error!")
	}

	if pa.Pid != 2766 {
		t.Error("pid parse error!")
	}
}

func Test_ProcessStateStopped(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:STOPPING pid:2766"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_STOPPED, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateStopped)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "STOPPING") {
		t.Error("from_state parse error!")
	}

	if pa.Pid != 2766 {
		t.Error("pid parse error!")
	}
}

func Test_ProcessStateUnknown(t *testing.T) {
	var tdata string = "processname:cat groupname:cat from_state:BACKOFF"
	p, err := protocol.Unmarshal(protocol.PROCESS_STATE_UNKNOWN, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessStateUnknown)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if !strings.EqualFold(pa.FromState, "BACKOFF") {
		t.Error("from_state parse error!")
	}
}

func Test_RemoteCommunication(t *testing.T) {
	var tdata string = "type:type\ndata"
	p, err := protocol.Unmarshal(protocol.REMOTE_COMMUNICATION, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.RemoteCommunication)
	if !strings.EqualFold(pa.Type, "type") {
		t.Error("type parse error!")
	}

	if !strings.EqualFold(pa.Data, "data") {
		t.Error("data parse error!")
	}
}

func Test_ProcessLogStdout(t *testing.T) {
	var tdata string = "processname:cat groupname:cat pid:2766\ndata"
	p, err := protocol.Unmarshal(protocol.PROCESS_LOG_STDOUT, tdata)
	if err != nil {
		t.Fatal(err)
	}

	pa := p.(*protocol.ProcessLogStdout)
	if !strings.EqualFold(pa.ProcessName, "cat") {
		t.Error("processname parse error!")
	}

	if !strings.EqualFold(pa.GroupName, "cat") {
		t.Error("groupname parse error!")
	}

	if pa.Pid != 2766 {
		t.Error("pid parse error!")
	}

	if !strings.EqualFold(pa.Data, "data") {
		t.Error("data parse error!")
	}
}
