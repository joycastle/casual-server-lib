package mysql

import (
	"fmt"
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
			SlowSqlTime: 1,
			SlowLogger:  "slow",
			StatLogger:  "stat",
		},

		"default-slave": MysqlConf{
			Addr:        "127.0.0.1",
			Username:    "root",
			Password:    "123456",
			Database:    "db_game",
			Options:     "charset=utf8mb4&parseTime=True",
			MaxIdle:     16,
			MaxOpen:     128,
			MaxLifeTime: time.Second * 300,
			SlowSqlTime: 1,
			SlowLogger:  "slow",
			StatLogger:  "stat",
		},
	}

	if err := InitMysql(configs); err != nil {
		t.Fatal(err)
	}

	if Get("default-master") == nil {
		t.Fatal()
	}

	fmt.Println(mysqlPoolMap)

	var rets []int64
	fmt.Println(Get("default-slave").Raw("SELECT * FROM `guild_help_request` WHERE time >= 1743326348 AND `guild_help_request`.`deleted_at` IS NULL ORDER BY id ASC LIMIT 1").Scan(&rets))

	time.Sleep(time.Second * 10)
}
