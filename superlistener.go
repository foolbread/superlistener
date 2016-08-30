//@auther foolbread
//@time 2016-08-29
//@file superlistener/superlistener.go
package main

import (
	"flag"

	"github.com/foolbread/superlistener/config"
	"github.com/foolbread/superlistener/listener"
	"github.com/foolbread/superlistener/log"

	"runtime"
)

var conf string

func init() {
	flag.StringVar(&conf, "f", "conf.ini", "conf.ini path!")
	flag.Parse()

	config.InitConfig(conf)
	log.InitLog(config.GetConfig().GetLogFile())
	listener.InitListener()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	select {}
}
