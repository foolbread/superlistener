//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/process_group_added.go
package protocol

import (
	"strings"
)

type ProcessGroupAdded struct {
	GroupName string
}

func unmarshalProcessGroupAdded(data string) *ProcessGroupAdded {
	gronamestrs := strings.Split(data, SPLIT_FIELD)
	ret := new(ProcessGroupAdded)
	ret.GroupName = gronamestrs[1]

	return ret
}
