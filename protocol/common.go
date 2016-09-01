//@auther foolbread
//@time 2016-08-30
//@file superlistener/protocol/common.go
package protocol

import (
	"errors"
)

const (
	EVENT                            = "EVENT"
	PROCESS_STATE                    = "PROCESS_STATE"
	PROCESS_STATE_STARTING           = "PROCESS_STATE_STARTING"
	PROCESS_STATE_RUNNING            = "PROCESS_STATE_RUNNING"
	PROCESS_STATE_BACKOFF            = "PROCESS_STATE_BACKOFF"
	PROCESS_STATE_STOPPING           = "PROCESS_STATE_STOPPING"
	PROCESS_STATE_EXITED             = "PROCESS_STATE_EXITED"
	PROCESS_STATE_STOPPED            = "PROCESS_STATE_STOPPED"
	PROCESS_STATE_FATAL              = "PROCESS_STATE_FATAL"
	PROCESS_STATE_UNKNOWN            = "PROCESS_STATE_UNKNOWN"
	REMOTE_COMMUNICATION             = "REMOTE_COMMUNICATION"
	PROCESS_LOG                      = "PROCESS_LOG"
	PROCESS_LOG_STDOUT               = "PROCESS_LOG_STDOUT"
	PROCESS_LOG_STDERR               = "PROCESS_LOG_STDERR"
	PROCESS_COMMUNICATION            = "PROCESS_COMMUNICATION"
	PROCESS_COMMUNICATION_STDOUT     = "PROCESS_COMMUNICATION_STDOUT"
	PROCESS_COMMUNICATION_STDERR     = "PROCESS_COMMUNICATION_STDERR"
	SUPERVISOR_STATE_CHANGE          = "SUPERVISOR_STATE_CHANGE"
	SUPERVISOR_STATE_CHANGE_RUNNING  = "SUPERVISOR_STATE_CHANGE_RUNNING"
	SUPERVISOR_STATE_CHANGE_STOPPING = "SUPERVISOR_STATE_CHANGE_STOPPING"
	TICK                             = "TICK"
	TICK_5                           = "TICK_5"
	TICK_60                          = "TICK_60"
	TICK_3600                        = "TICK_3600"
	PROCESS_GROUP                    = "PROCESS_GROUP"
	PROCESS_GROUP_ADDED              = "PROCESS_GROUP_ADDED"
	PROCESS_GROUP_REMOVED            = "PROCESS_GROUP_REMOVED"
)

const (
	HEAD_FIELD_CNT                   = 7
	PROCESS_STATE_FATAL_FIELD_CNT    = 3
	PROCESS_STATE_BACKOFF_FIELD_CNT  = 4
	PROCESS_STATE_STARTING_FIELD_CNT = 4
)

func Unmarshal(event string, data string) (interface{}, error) {
	var ret interface{} = nil
	var err error = nil

	switch event {
	case PROCESS_STATE_STARTING:
		ret = unmarshalProcessStateStarting(data)
	case PROCESS_STATE_BACKOFF:
		ret = unmarshalProcessStateBackoff(data)
	case PROCESS_STATE_FATAL:
		ret = unmarshalProcessStateFatal(data)
	}

	if ret == nil {
		err = errors.New("unknow data,parse fail!")
	}

	return ret, err
}
