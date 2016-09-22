//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/process_group_removed.go
package protocol

import (
	"strings"
)

type ProcessGroupRemoved struct {
	GroupName string
}

func unmarshalProcessGroupRemoved(data string) *ProcessGroupRemoved {
	gronamestrs := strings.Split(data, SPLIT_FIELD)
	ret := new(ProcessGroupRemoved)
	ret.GroupName = gronamestrs[1]

	return ret
}
