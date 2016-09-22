//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/tick_3600.go
package protocol

import (
	"strconv"
	"strings"
)

type Tick3600 struct {
	When int
}

func unmarshalTick3600(data string) *Tick3600 {
	whenstrs := strings.Split(data, SPLIT_FIELD)
	ret := new(Tick3600)
	ret.When, _ = strconv.Atoi(whenstrs[1])

	return ret
}
