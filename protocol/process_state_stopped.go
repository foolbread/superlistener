//@auther foolbread
//@time 2016-09-01
//@file superlistener/protocol/process_state_stopped.go
package protocol

import (
	"strconv"
	"strings"
)

type ProcessStateStopped struct {
	ProcessName string
	GroupName   string
	FromState   string
	Pid         int
}

func unmarshalProcessStateStopped(data string) *ProcessStateStopped {
	fields := strings.Split(data, SPLIT_LINE)
	if len(fields) != PROCESS_STATE_STOPPED_FIELD_CNT {
		return nil
	}

	ret := new(ProcessStateStopped)
	pronamestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.ProcessName = pronamestrs[1]
	gronamestrs := strings.Split(fields[1], SPLIT_FIELD)
	ret.GroupName = gronamestrs[1]
	fromstastrs := strings.Split(fields[2], SPLIT_FIELD)
	ret.FromState = fromstastrs[1]
	pidstrs := strings.Split(fields[3], SPLIT_FIELD)
	ret.Pid, _ = strconv.Atoi(pidstrs[1])

	return ret
}
