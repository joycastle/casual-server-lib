package config

import (
	"fmt"
	"time"

	"github.com/joycastle/casual-server-lib/mysql"
	"github.com/spf13/viper"
)

const (
	CFG_MYSQL             = "s-mysql"
	CFG_MYSQL_ADDR        = "addr"
	CFG_MYSQL_USERNAME    = "username"
	CFG_MYSQL_PASSWORD    = "password"
	CFG_MYSQL_DATABASE    = "database"
	CFG_MYSQL_OPTIONS     = "options"
	CFG_MYSQL_MAXIDEL     = "maxidle"
	CFG_MYSQL_MAXOPEN     = "maxopen"
	CFG_MYSQL_MAXLIFETIME = "maxlifetime"
	CFG_MYSQL_SLOWLOGTIME = "slowsqltime"
	CFG_MYSQL_SLOWLOG     = "slowlog"
	CFG_MYSQL_STATLOG     = "statlog"
)

var Mysql map[string]mysql.MysqlConf = make(map[string]mysql.MysqlConf)

func init() {
	RegisterParser(parseMysql)
}

func parseMysql(v *viper.Viper) error {
	mps := v.GetStringMap(CFG_MYSQL)

	if len(mps) == 0 {
		return ErrFileNotExists
	}

	for k, v := range mps {
		vv := v.(map[string]interface{})

		c := mysql.MysqlConf{}

		if s, ok := vv[CFG_MYSQL_ADDR]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_ADDR)
		} else {
			c.Addr = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_USERNAME]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_USERNAME)
		} else {
			c.Username = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_PASSWORD]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_PASSWORD)
		} else {
			c.Password = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_DATABASE]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_DATABASE)
		} else {
			c.Database = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_OPTIONS]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_OPTIONS)
		} else {
			c.Options = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_MAXIDEL]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_MAXIDEL)
		} else {
			c.MaxIdle = s.(int)
		}

		if s, ok := vv[CFG_MYSQL_MAXOPEN]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_MAXOPEN)
		} else {
			c.MaxOpen = s.(int)
		}

		if s, ok := vv[CFG_MYSQL_MAXLIFETIME]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_MAXLIFETIME)
		} else {
			c.MaxLifeTime, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_MYSQL_SLOWLOGTIME]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_SLOWLOGTIME)
		} else {
			c.SlowSqlTime, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_MYSQL_SLOWLOG]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_SLOWLOG)
		} else {
			c.SlowLogger = s.(string)
		}

		if s, ok := vv[CFG_MYSQL_STATLOG]; !ok {
			return fmt.Errorf("MYSQL config file not contains \"%s\"", CFG_MYSQL_STATLOG)
		} else {
			c.StatLogger = s.(string)
		}

		Mysql[k] = c
	}

	return nil
}
