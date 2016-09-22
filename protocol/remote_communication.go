//@auther foolbread
//@time 2016-09-01
//@file superlistener/protocol/remote_communication.go
package protocol

import (
	"strings"
)

type RemoteCommunication struct {
	Type string
	Data string
}

func unmarshalRemoteCommunication(data string) *RemoteCommunication {
	fields := strings.Split(data, SPLIT_MUTILLINE)
	if len(fields) != PROCESS_REMOTE_COMMUNICATION_LINE {
		return nil
	}

	ret := new(RemoteCommunication)
	typestrs := strings.Split(fields[0], SPLIT_FIELD)
	ret.Type = typestrs[1]
	ret.Data = fields[1]

	return ret
}
