package persistence

import (
	"blog.io/config"
	"github.com/garyburd/redigo/redis"
	"time"
)

// redis 连接池
var pool *redis.Pool

// 根据配置初始化打开redis连接
func init() {
	conf := config.Config().RedisCfg
	pool := &redis.Pool{
		MaxIdle: 20,
		MaxActive: 30,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Host + ":" + conf.Port)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedisPool() *redis.Pool {
	return pool
}