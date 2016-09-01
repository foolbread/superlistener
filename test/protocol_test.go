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
