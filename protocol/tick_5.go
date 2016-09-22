//@auther foolbread
//@time 2016-09-22
//@file superlistener/protocol/tick_5.go
package protocol

import (
	"strconv"
	"strings"
)

type Tick5 struct {
	When int
}

func unmarshalTick5(data string) *Tick5 {
	whenstrs := strings.Split(data, SPLIT_FIELD)
	ret := new(Tick5)
	ret.When, _ = strconv.Atoi(whenstrs[1])

	return ret
}
