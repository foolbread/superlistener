//@auther foolbread
//@time 2016-08-30
//@file superlistener/protocol/process_state_fatal.go
package protocol

import (
	"strings"
)

type ProcessStateFatal struct {
	ProcessName string
	GroupName   string
	FromState   string
}

func UnmarshalProcessStateFatal(data string) *ProcessStateFatal {
	fields := strings.Split(data, SPLIT_LINE)
	if len(fields) != PROCESS_STATE_FATAL_FIELD_CNT {
		return nil
	}

	ret := new(ProcessStateFatal)
	pronamestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.ProcessName = pronamestrs[1]
	gronamestrs := strings.Split(fields[1], SPLIT_FIELD)
	ret.GroupName = gronamestrs[1]
	fromstastrs := strings.Split(fields[2], SPLIT_FIELD)
	ret.FromState = fromstastrs[1]

	return ret
}
