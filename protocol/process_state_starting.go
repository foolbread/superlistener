//@auther foolbread
//@time 2016-09-01
//@file superlistener/protocol/process_state_starting.go
package protocol

import (
	"strconv"
	"strings"
)

type ProcessStateStarting struct {
	ProcessName string
	GroupName   string
	FromState   string
	Tries       int
}

func unmarshalProcessStateStarting(data string) *ProcessStateStarting {
	fields := strings.Split(data, SPLIT_LINE)
	if len(fields) != PROCESS_STATE_STARTING_FIELD_CNT {
		return nil
	}

	ret := new(ProcessStateStarting)
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
