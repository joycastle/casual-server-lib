package redis

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/joycastle/casual-server-lib/log"
)

var (
	redisPoolMap map[string]*redis.Pool
	redisNodes   []string
)

type RedisConf struct {
	Addr           string
	Password       string
	MaxActive      int
	MaxIdle        int
	IdleTimeout    time.Duration
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	TestInterval   time.Duration

	ErrLogger  string
	StatLogger string
}

func GetConn(n string) redis.Conn {
	if pool, ok := redisPoolMap[n]; ok {
		return pool.Get()
	}
	log.Get("error").Fatalf(fmt.Sprintf("Redis node \"%s\" not exists, choose from %v", n, redisNodes))
	panic(fmt.Sprintf("Redis node \"%s\" not exists, choose from %v", n, redisNodes))
	return nil
}

func InitRedis(configs map[string]RedisConf) {
	redisPoolMap = make(map[string]*redis.Pool, len(configs))

	for sn, config := range configs {
		redisPoolMap[sn] = GetRedisConn(sn, config)
		redisNodes = append(redisNodes, sn)
	}
}

func GetRedisConn(sn string, config RedisConf) *redis.Pool {

	rp := &redis.Pool{
		MaxActive:   config.MaxActive,
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {

			var (
				conn redis.Conn
				err  error
			)

			addrs := strings.Split(config.Addr, ",")
			addr := addrs[rand.Intn(len(addrs))]
			conn, err = redis.DialTimeout("tcp", addr, config.ConnectTimeout, config.ReadTimeout, config.WriteTimeout)

			if err != nil {
				log.Get(config.ErrLogger).Warnf("Redis connect failed. %s, addr:%s", err, addr)
				return nil, err
			}

			if config.Password != "" {
				if _, err := conn.Do("AUTH", config.Password); err != nil {
					log.Get(config.ErrLogger).Warnf("Redis auth failed. %s, addr:%s", err, addr)
					conn.Close()
					return nil, err
				}
			}

			if _, err = conn.Do("PING"); err != nil {
				log.Get(config.ErrLogger).Warnf("Redis ping failed. %s, addr:%s", err, addr)
				return nil, err
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < config.TestInterval {
				return nil
			}
			if _, err := conn.Do("PING"); err != nil {
				log.Get(config.ErrLogger).Warnf("Redis TestOnBorrow failed. %s", err)
				return err
			}
			return nil
		},
	}

	go func() {
		// PoolStats contains pool statistics.
		//type PoolStats struct {
		// ActiveCount is the number of connections in the pool. The count includes
		// idle connections and connections in use.
		// ActiveCount int
		// IdleCount is the number of idle connections in the pool.
		//IdleCount int
		//}
		var lastInfs string

		for {
			stat := rp.Stats()
			infs := fmt.Sprintf("Redis Pool ActiveCount:%d, IdleCount:%d node:%s", stat.ActiveCount, stat.IdleCount, sn)
			if infs != lastInfs {
				log.Get(config.StatLogger).Info(infs)
				lastInfs = infs
			}

			time.Sleep(time.Second * 10)
		}
	}()

	return rp
}
