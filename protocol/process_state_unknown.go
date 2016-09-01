//@auther foolbread
//@time 2016-09-01
//@file superlistener/protocol/process_state_unknown.go
package protocol

import (
	"strings"
)

type ProcessStateUnknown struct {
	ProcessName string
	GroupName   string
	FromState   string
}

func unmarshalProcessStateUnknown(data string) *ProcessStateUnknown {
	fields := strings.Split(data, SPLIT_LINE)
	if len(fields) != PROCESS_STATE_UNKNOWN_FIELD_CNT {
		return nil
	}

	ret := new(ProcessStateUnknown)
	pronamestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.ProcessName = pronamestrs[1]
	gronamestrs := strings.Split(fields[1], SPLIT_FIELD)
	ret.GroupName = gronamestrs[1]
	fromstastrs := strings.Split(fields[2], SPLIT_FIELD)
	ret.FromState = fromstastrs[1]

	return ret
}
