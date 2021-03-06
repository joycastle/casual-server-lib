package config

import (
	"testing"
	"time"
)

func Test_config(t *testing.T) {
	if err := InitConfig("./example.yaml"); err != nil {
		t.Fatal(err)
	}

	if Logs["main"].Output != "./main.log-*-*-*" || Logs["main"].Level != "INFO" || Logs["main"].TraceOffset != 0 || Logs["error"].TraceOffset != 10 {
		t.Fatal("parse error")
	}

	if Redis["main"].Addr != "127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379" || Redis["main"].Password != "123456" {
		t.Fatal("Redis", Redis)
	}

	if Redis["default"].MaxActive != 128 || Redis["default"].MaxIdle != 16 || Redis["default"].ConnectTimeout != time.Millisecond*500 {
		t.Fatal("Redis", Redis)
	}

	if Mysql["default-master"].Addr != "127.0.0.1" || Mysql["default-master"].Username != "root" || Mysql["default-slave"].Password != "123456" || Mysql["default-master"].Database != "db_game" {
		t.Fatal("Mysql", Mysql)
	}

	if Mysql["default-master"].MaxIdle != 16 || Mysql["default-master"].MaxOpen != 128 || Mysql["default-slave"].MaxLifeTime != 5*time.Minute || Mysql["default-master"].SlowLogger != "slow" {
		t.Fatal("Mysql", Mysql)
	}

	if Grpc["default"] != "127.0.0.1:9002" || Grpc["chat"] != "127.0.0.1:9001" {
		t.Fatal("GRPC", Grpc)
	}
}

func Test_config_dir(t *testing.T) {
	if err := InitConfig("./conf_dir"); err != nil {
		t.Fatal(err)
	}

	if Logs["main"].Output != "./main.log-*-*-*" || Logs["main"].Level != "INFO" || Logs["main"].TraceOffset != 0 || Logs["error"].TraceOffset != 10 {
		t.Fatal("parse error")
	}

	if Redis["main"].Addr != "127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379" || Redis["main"].Password != "123456" {
		t.Fatal("Redis", Redis)
	}

	if Redis["default"].MaxActive != 128 || Redis["default"].MaxIdle != 16 || Redis["default"].ConnectTimeout != time.Millisecond*500 {
		t.Fatal("Redis", Redis)
	}

	if Mysql["default-master"].Addr != "127.0.0.1" || Mysql["default-master"].Username != "root" || Mysql["default-slave"].Password != "123456" || Mysql["default-master"].Database != "db_game" {
		t.Fatal("Mysql", Mysql)
	}

	if Mysql["default-master"].MaxIdle != 16 || Mysql["default-master"].MaxOpen != 128 || Mysql["default-slave"].MaxLifeTime != 5*time.Minute || Mysql["default-master"].SlowLogger != "slow" {
		t.Fatal("Mysql", Mysql)
	}
}
