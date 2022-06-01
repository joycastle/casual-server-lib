package redis

import (
	"testing"
	"time"
)

func init() {
	configs := make(map[string]RedisConf)
	configs["default"] = RedisConf{
		Addr:           "127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379",
		Password:       "123456",
		MaxActive:      32,
		MaxIdle:        16,
		IdleTimeout:    time.Second * 1800,
		ConnectTimeout: time.Second * 10,
		ReadTimeout:    time.Second * 2,
		WriteTimeout:   time.Second * 2,
		TestInterval:   time.Second * 300,
	}

	configs["new"] = RedisConf{
		Addr:           "127.0.0.1:26379,127.0.0.1:26379,127.0.0.1:26379",
		Password:       "123456",
		MaxActive:      32,
		MaxIdle:        16,
		IdleTimeout:    time.Second * 1800,
		ConnectTimeout: time.Second * 10,
		ReadTimeout:    time.Second * 2,
		WriteTimeout:   time.Second * 2,
		TestInterval:   time.Second * 300,
	}

	InitRedis(configs)
}

func Test_Redis(t *testing.T) {
	if b, err := GetBytes("new", "levin"); err.Error() != "dial tcp 127.0.0.1:26379: connect: connection refused" {
		t.Fatal(b, err)
	}

	if r, err := Del("default", "levin"); r > 1 || err != nil {
		t.Fatal(r, err)
	}

	if b, err := GetBytes("default", "levin"); err == nil {
		t.Fatal(b, err)
	}

	if r, err := Set("default", "levin", "hello"); r != "OK" || err != nil {
		t.Fatal(r, err)
	}

	if r, err := GetInt("default", "levin"); err == nil {
		t.Fatal(r, err)
	}

	if r, err := Set("default", "levin", 123456); r != "OK" || err != nil {
		t.Fatal(r, err)
	}

	if r, err := GetInt("default", "levin"); r != 123456 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := GetInt64("default", "levin"); r != 123456 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := Incr("default", "levin"); r != 123457 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := IncrBy("default", "levin", 10); r != 123467 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := DecrBy("default", "levin", 10); r != 123457 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := Decr("default", "levin"); r != 123456 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := SetEx("default", "levin", "Hi", 10); r != "OK" || err != nil {
		t.Fatal(r, err)
	}

	time.Sleep(time.Second)

	if r, err := TTL("default", "levin"); r != 9 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := Expire("default", "levin", 2); r != 1 || err != nil {
		t.Fatal(r, err)
	}

	time.Sleep(time.Second)

	if r, err := TTL("default", "levin"); r != 1 || err != nil {
		t.Fatal(r, err)
	}

	if r, err := Exists("default", "levin"); r != 1 || err != nil {
		t.Fatal(r, err)
	}

	time.Sleep(time.Second * 2)

	if r, err := Exists("default", "levin"); r != 0 || err != nil {
		t.Fatal(r, err)
	}
}
