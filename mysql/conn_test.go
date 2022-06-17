package mysql

import (
	"testing"
	"time"
)

func Test_Mysql(t *testing.T) {
	configs := map[string]MysqlConf{
		"default-master": MysqlConf{
			Addr:        "127.0.0.1",
			Username:    "root",
			Password:    "123456",
			Database:    "db_game",
			Options:     "charset=utf8mb4&parseTime=True",
			MaxIdle:     16,
			MaxOpen:     128,
			MaxLifeTime: time.Second * 300,
			SlowSqlTime: 0,
			SlowLogger:  "slow",
			ErrLogger:   "error",
			StatLogger:  "stat",
		},
	}

	if err := InitMysql(configs); err != nil {
		t.Fatal(err)
	}

	if Get("default-master") == nil {
		t.Fatal()
	}
}
