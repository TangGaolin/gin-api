package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

var RedisPool *redis.Pool

func Init() {
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	password := viper.GetString("redis.password")
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedisDataByHashKey(hashKey string, db int) (map[string]string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	conn.Do("SELECT", db)
	hashValue, err := redis.StringMap(conn.Do("HGETALL", hashKey))
	if err != nil {
		return nil, err
	}
	return hashValue, nil
}

func GetRedisDataByHashKeyValue(hashKey string, valueKey string, db int) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	conn.Do("SELECT", db)
	value, err := redis.String(conn.Do("HGET", hashKey, valueKey))
	if err != nil {
		return "", err
	}
	return value, nil
}

func SetRedisDataByHashKey(hashKey string, valueKey string, value []byte, db int) error {
	conn := RedisPool.Get()
	defer conn.Close()
	conn.Do("SELECT", db)
	_, err := conn.Do("HSET", hashKey, valueKey, value)
	if err != nil {
		return err
	}
	return nil
}

func DelRedisDataByHashKey(hashKey string, valueKey string, db int) error {
	conn := RedisPool.Get()
	defer conn.Close()
	conn.Do("SELECT", db)
	_, err := conn.Do("HDEL", hashKey, valueKey)
	if err != nil {
		return err
	}
	return nil
}
