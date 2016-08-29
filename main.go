//@auther foolbread
//@time 2016-08-29
//@file superlistener/main.go
package main

import (
	"flag"
	"superlistener/config"
	"superlistener/log"
)

var conf string

func init() {
	flag.StringVar(&conf, "f", "conf.ini", "conf.ini path!")
	flag.Parse()

	config.InitConfig(conf)
	log.InitLog(config.GetConfig().GetLogFile())
}

func main() {

}
