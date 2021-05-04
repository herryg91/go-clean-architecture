package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisConfig struct {
	Host            string
	Port            string
	Password        string
	MaxIdle         int
	MaxActive       int
	MaxIdleTimeOut  int
	MaxConnLifetime int
	Wait            bool
}
type redisOption func(*redisConfig)

func Connect(host string, port string, password string, options ...redisOption) (*redis.Pool, error) {
	db := &redisConfig{
		Host:            host,
		Port:            port,
		Password:        password,
		MaxIdle:         5,
		MaxActive:       10,
		MaxIdleTimeOut:  5,
		MaxConnLifetime: 10,
		Wait:            true,
	}
	if db.Port == "" {
		db.Port = "6379"
	}

	for _, o := range options {
		o(db)
	}

	return connect(db)

}

func connect(param *redisConfig) (*redis.Pool, error) {
	redisPool := &redis.Pool{
		MaxIdle:         param.MaxIdle,
		MaxActive:       param.MaxActive,
		Wait:            param.Wait,
		MaxConnLifetime: time.Duration(param.MaxConnLifetime) * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", param.Host, param.Port))
			if err != nil {
				return nil, err
			}
			if len(param.Password) > 0 {
				if _, err := c.Do("AUTH", param.Password); err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}

	return redisPool, nil
}

func SetMaxIdle(conns int) redisOption {
	return func(c *redisConfig) {
		if conns > 0 {
			c.MaxIdle = conns
		}
	}
}

func SetMaxIdleTimeout(timeout int) redisOption {
	return func(c *redisConfig) {
		if timeout > 0 {
			c.MaxIdleTimeOut = timeout
		}
	}
}

func SetMaxActive(conns int) redisOption {
	return func(c *redisConfig) {
		if conns > 0 {
			c.MaxActive = conns
		}
	}
}

func SetConnMaxLifetime(timeout int) redisOption {
	return func(c *redisConfig) {
		if timeout > 0 {
			c.MaxConnLifetime = timeout
		}
	}
}
