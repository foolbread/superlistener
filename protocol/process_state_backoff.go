//@auther foolbread
//@time 2016-08-30
//@file superlistener/protocol/process_state_backoff.go
package protocol

import (
	"strconv"
	"strings"
)

type ProcessStateBackoff struct {
	ProcessName string
	GroupName   string
	FromState   string
	Tries       int
}

func unmarshalProcessStateBackoff(data string) *ProcessStateBackoff {
	fields := strings.Split(data, SPLIT_LINE)
	if len(fields) != PROCESS_STATE_BACKOFF_FIELD_CNT {
		return nil
	}

	ret := new(ProcessStateBackoff)
	pronamestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.ProcessName = pronamestrs[1]
	gronamestrs := strings.Split(fields[1], SPLIT_FIELD)
	ret.GroupName = gronamestrs[1]
	fromstastrs := strings.Split(fields[2], SPLIT_FIELD)
	ret.FromState = fromstastrs[1]
	triesstrs := strings.Split(fields[3], SPLIT_FIELD)
	ret.Tries, _ = strconv.Atoi(triesstrs[1])

	return ret
}
