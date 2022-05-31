package config

import (
	"fmt"

	"github.com/joycastle/casual-server-lib/log"
	"github.com/spf13/viper"
)

const (
	CFG_LOG        = "logs"
	CFG_LOG_OUTPUT = "output"
	CFG_LOG_LEVEL  = "level"
)

var Logs map[string]log.LogConf = make(map[string]log.LogConf)

func init() {
	RegisterParser(parseLog)
}

func parseLog(v *viper.Viper) error {
	mps := v.GetStringMap(CFG_LOG)

	for k, v := range mps {
		vv := v.(map[string]interface{})

		c := log.LogConf{}

		if s, ok := vv[CFG_LOG_OUTPUT]; !ok {
			return fmt.Errorf("LOG config file not contains \"%s\"", CFG_LOG_OUTPUT)
		} else {
			c.Output = s.(string)
		}

		if s, ok := vv[CFG_LOG_LEVEL]; !ok {
			return fmt.Errorf("LOG config file not contains \"%s\"", CFG_LOG_LEVEL)
		} else {
			c.Level = s.(string)
		}

		Logs[k] = c
	}

	return nil
}