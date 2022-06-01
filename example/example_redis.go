package main

import (
	"runtime"
	"time"

	"github.com/joycastle/casual-server-lib/log"
	"github.com/joycastle/casual-server-lib/redis"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	redisConfigs := map[string]redis.RedisConf{
		"default": redis.RedisConf{
			Addr:           "127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379",
			Password:       "123456",
			MaxActive:      32,
			MaxIdle:        16,
			IdleTimeout:    time.Second * 1800,
			ConnectTimeout: time.Second * 10,
			ReadTimeout:    time.Second * 2,
			WriteTimeout:   time.Second * 2,
			TestInterval:   time.Second * 300,
		},

		"product": redis.RedisConf{
			Addr:           "127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379",
			Password:       "123456",
			MaxActive:      32,
			MaxIdle:        16,
			IdleTimeout:    time.Second * 1800,
			ConnectTimeout: time.Second * 10,
			ReadTimeout:    time.Second * 2,
			WriteTimeout:   time.Second * 2,
			TestInterval:   time.Second * 300,
		},
	}

	//init redis log
	redis.SetLogger(log.DefaultLogger)

	//init redis
	redis.InitRedis(redisConfigs)

	//first: node name eg: default or product
	//second: redis key
	if r, err := redis.Del("default", "levin"); r > 1 || err != nil {
		log.Fatal(r, err)
	}

	if b, err := redis.GetBytes("default", "levin"); err == nil {
		log.Fatal(b, err)
	}

	if r, err := redis.Set("default", "levin", "hello"); r != "OK" || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.GetInt("default", "levin"); err == nil {
		log.Fatal(r, err)
	}

	if r, err := redis.Set("default", "levin", 123456); r != "OK" || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.GetInt("default", "levin"); r != 123456 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.GetInt64("default", "levin"); r != 123456 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.Incr("default", "levin"); r != 123457 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.IncrBy("default", "levin", 10); r != 123467 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.DecrBy("default", "levin", 10); r != 123457 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.Decr("default", "levin"); r != 123456 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.SetEx("default", "levin", "Hi", 10); r != "OK" || err != nil {
		log.Fatal(r, err)
	}

	time.Sleep(time.Second)

	if r, err := redis.TTL("default", "levin"); r != 9 || err != nil {
		log.Fatal(r, err)
	}

	if r, err := redis.Expire("default", "levin", 2); r != 1 || err != nil {
		log.Fatal(r, err)
	}

	//......
	//......
}
