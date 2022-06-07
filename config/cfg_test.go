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
}
