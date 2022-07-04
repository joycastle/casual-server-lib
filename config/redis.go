package config

import (
	"fmt"
	"time"

	"github.com/joycastle/casual-server-lib/redis"
	"github.com/spf13/viper"
)

const (
	CFG_REDIS                 = "s-redis"
	CFG_REDIS_ADDR            = "addr"
	CFG_REDIS_PASSWORD        = "password"
	CFG_REDIS_MAX_ACTIVE      = "maxactive"
	CFG_REDIS_MAX_IDEL        = "maxidle"
	CFG_REDIS_IDEL_TIMEOUT    = "idletimeout"
	CFG_REDIS_CONNECT_TIMEOUT = "connecttimeout"
	CFG_REDIS_READ_TIMEOUT    = "readtimeout"
	CFG_REDIS_WRITE_TIMEOUT   = "writetimeout"
	CFG_REDIS_TEST_INTERVAL   = "testinterval"

	CFG_REDIS_ERRORLOG = "errorlog"
	CFG_REDIS_STATLOG  = "statlog"
)

var Redis map[string]redis.RedisConf = make(map[string]redis.RedisConf)

func init() {
	RegisterParser(parseRedis)
}

func parseRedis(v *viper.Viper) error {
	mps := v.GetStringMap(CFG_REDIS)

	if len(mps) == 0 {
		return ErrFileNotExists
	}

	for k, v := range mps {
		vv := v.(map[string]interface{})

		c := redis.RedisConf{}

		if s, ok := vv[CFG_REDIS_ADDR]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_ADDR)
		} else {
			c.Addr = s.(string)
		}

		if s, ok := vv[CFG_REDIS_PASSWORD]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_PASSWORD)
		} else {
			c.Password = s.(string)
		}

		if s, ok := vv[CFG_REDIS_MAX_ACTIVE]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_MAX_ACTIVE)
		} else {
			c.MaxActive = s.(int)
		}

		if s, ok := vv[CFG_REDIS_MAX_IDEL]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_MAX_IDEL)
		} else {
			c.MaxIdle = s.(int)
		}

		if s, ok := vv[CFG_REDIS_IDEL_TIMEOUT]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_IDEL_TIMEOUT)
		} else {
			c.IdleTimeout, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_REDIS_CONNECT_TIMEOUT]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_CONNECT_TIMEOUT)
		} else {
			c.ConnectTimeout, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_REDIS_READ_TIMEOUT]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_READ_TIMEOUT)
		} else {
			c.ReadTimeout, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_REDIS_WRITE_TIMEOUT]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_WRITE_TIMEOUT)
		} else {
			c.WriteTimeout, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_REDIS_TEST_INTERVAL]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_TEST_INTERVAL)
		} else {
			c.TestInterval, _ = time.ParseDuration(s.(string))
		}

		if s, ok := vv[CFG_REDIS_ERRORLOG]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_ERRORLOG)
		} else {
			c.ErrLogger = s.(string)
		}

		if s, ok := vv[CFG_REDIS_STATLOG]; !ok {
			return fmt.Errorf("REDIS config file not contains \"%s\"", CFG_REDIS_STATLOG)
		} else {
			c.StatLogger = s.(string)
		}

		Redis[k] = c
	}

	return nil
}
