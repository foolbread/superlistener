//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/tick_60.go
package protocol

import (
	"strconv"
	"strings"
)

type Tick60 struct {
	When int
}

func unmarshalTick60(data string) *Tick60 {
	whenstrs := strings.Split(data, SPLIT_FIELD)
	ret := new(Tick60)
	ret.When, _ = strconv.Atoi(whenstrs[1])

	return ret
}
