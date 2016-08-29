//@auther foolbread
//@time 2016-08-29
//@file superlistener/config/config.go
package config

import (
	"github.com/foolbread/fbcommon/config"
	"github.com/foolbread/fbcommon/golog"
)

func InitConfig(file string) {
	conf, err := config.LoadConfigByFile(file)
	if err != nil {
		golog.Critical(err)
	}

	val := conf.MustString("log", "log_file", "")
	if len(val) == 0 {
		golog.Critical(err)
	}
	g_config.setLogFile(val)

	val = conf.MustString("listener", "listener_file", "")
	if len(val) == 0 {
		golog.Critical(err)
	}
	g_config.setListenerFile(val)
}

func GetConfig() *superListenerConfig {
	return g_config
}

var g_config *superListenerConfig = new(superListenerConfig)

type superListenerConfig struct {
	logFile      string
	listenerFile string
}

func (c *superListenerConfig) setLogFile(s string) {
	c.logFile = s
}

func (c *superListenerConfig) setListenerFile(s string) {
	c.listenerFile = s
}

func (c *superListenerConfig) GetLogFile() string {
	return c.logFile
}

func (c *superListenerConfig) GetListenerFile() string {
	return c.listenerFile
}
