//@auther foolbread
//@time 2016-08-30
//@file superlistener/protocol/head.go
package protocol

import (
	"strconv"
	"strings"
)

const (
	SPLIT_LINE      = " "
	SPLIT_FIELD     = ":"
	SPLIT_MUTILLINE = "\n"
)

type supervisorHead struct {
	Version     string
	Server      string
	Serial      int
	Pool        string
	Poolserial  int
	Eventname   string
	Playloadlen int
}

func UnmarshalSupervisorHead(line string) *supervisorHead {
	fields := strings.Split(line, SPLIT_LINE)
	if len(fields) != 7 {
		return nil
	}

	ret := new(supervisorHead)
	verstr := strings.Split(fields[0], SPLIT_FIELD)
	ret.Version = verstr[1]
	serstr := strings.Split(fields[1], SPLIT_FIELD)
	ret.Server = serstr[1]
	serialstr := strings.Split(fields[2], SPLIT_FIELD)
	ret.Serial, _ = strconv.Atoi(serialstr[1])
	poolstr := strings.Split(fields[3], SPLIT_FIELD)
	ret.Pool = poolstr[1]
	pserialstr := strings.Split(fields[4], SPLIT_FIELD)
	ret.Poolserial, _ = strconv.Atoi(pserialstr[1])
	eventstr := strings.Split(fields[5], SPLIT_FIELD)
	ret.Eventname = eventstr[1]
	lenstr := strings.Split(fields[6], SPLIT_FIELD)
	ret.Playloadlen, _ = strconv.Atoi(lenstr[1])

	return ret
}
