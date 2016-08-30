//@auther foolbread
//@time 2016-08-29
//@file superlistener/log/log.go
package log

import (
	"github.com/cihub/seelog"
	"github.com/foolbread/fbcommon/golog"
)

func InitLog(log string) {
	var err error
	g_log, err = seelog.LoggerFromConfigAsFile(log)
	if err != nil {
		golog.Critical(err)
	}
}

func LogFile() seelog.LoggerInterface {
	return g_log
}

var g_log seelog.LoggerInterface
