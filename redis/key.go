package redis

import (
	"github.com/garyburd/redigo/redis"
)

func GetBytes(sn string, key string) ([]byte, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("GET", key))
	return b, err
}

func GetString(sn string, key string) (string, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.String(conn.Do("GET", key))
	return b, err
}

func GetInt(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("GET", key))
	return b, err
}

func GetInt64(sn string, key string) (int64, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int64(conn.Do("GET", key))
	return b, err
}

func Set(sn string, key string, v any) (string, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.String(conn.Do("SET", key, v))
	return b, err
}

func SetEx(sn string, key string, v any, ex int) (string, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.String(conn.Do("SETEX", key, ex, v))
	return b, err
}

func TTL(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("TTL", key))
	return b, err
}

func Expire(sn string, key string, v int) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("EXPIRE", key, v))
	return b, err
}

func Exists(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("EXISTS", key))
	return b, err
}

func Del(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("DEL", key))
	return b, err
}

func Incr(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("INCR", key))
	return b, err
}

func IncrBy(sn string, key string, step int) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("INCRBY", key, step))
	return b, err
}

func Decr(sn string, key string) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("DECR", key))
	return b, err
}

func DecrBy(sn string, key string, step int) (int, error) {
	conn := GetConn(sn)
	defer conn.Close()
	b, err := redis.Int(conn.Do("DECRBY", key, step))
	return b, err
}
