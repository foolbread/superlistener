//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/process_communication_stdout.go
package protocol

import (
	"strconv"
	"strings"
)

type ProcessCommunicationStdout struct {
	ProcessName string
	GroupName   string
	Pid         int
	Data        string
}

func unmarshalProcessCommunicationStdout(data string) *ProcessCommunicationStdout {
	lines := strings.Split(data, SPLIT_MUTILLINE)
	if len(lines) != PROCESS_COMMUNICATION_STDOUT_LINE {
		return nil
	}

	fields := strings.Split(lines[0], SPLIT_LINE)
	if len(fields) != PROCESS_COMMUNICATION_STDOUT_CNT {
		return nil
	}

	ret := new(ProcessCommunicationStdout)
	pronamestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.ProcessName = pronamestrs[1]
	gronamestrs := strings.Split(fields[1], SPLIT_FIELD)
	ret.GroupName = gronamestrs[1]
	pidstrs := strings.Split(fields[2], SPLIT_FIELD)
	ret.Pid, _ = strconv.Atoi(pidstrs[1])
	ret.Data = lines[1]

	return ret
}
