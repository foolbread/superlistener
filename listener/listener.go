//@auther foolbread
//@time 2016-08-29
//@file superlistener/listener/listener.go
package listener

import (
	"bufio"
	"fmt"
	"os"

	"github.com/foolbread/superlistener/config"
	"github.com/foolbread/superlistener/log"
	"github.com/foolbread/superlistener/method"
	"github.com/foolbread/superlistener/protocol"
)

const (
	READY_STR  = "READY\n"
	FINISH_STR = "RESULT 2\nOK"
)

func InitListener() {
	infos := config.GetConfig().GetProgramInfos()
	for _, v := range infos.Programs {
		switch v.Method {
		//define by yourself
		}
	}
}

var g_listener *supervisorListener = newSupervisorListener()

type supervisorListener struct {
	eventmap map[string][]method.NotifyMethod
}

func newSupervisorListener() *supervisorListener {
	ret := new(supervisorListener)
	ret.eventmap = make(map[string][]method.NotifyMethod)

	return ret
}

func (l *supervisorListener) run() {

	var data [4096]byte
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(READY_STR)

		line, _, err := rd.ReadLine()
		if err != nil {
			log.LogFile().Error(err)
			os.Exit(1)
		}

		head := protocol.UnmarshalSupervisorHead(string(line))
		if head == nil {
			log.LogFile().Error("parse head error,line:", string(line))
			os.Exit(1)
		}

		_, err = rd.Read(data[:head.Playloadlen])
		if err != nil {
			log.LogFile().Error(err)
			os.Exit(1)
		}

		e, ok := l.eventmap[head.Eventname]
		if ok {
			for _, v := range e {
				go v.DoNotify(head.Eventname, data[:head.Playloadlen])
			}
		}

		log.LogFile().Info(string(data[:head.Playloadlen]))

		fmt.Print(FINISH_STR)
	}
}
