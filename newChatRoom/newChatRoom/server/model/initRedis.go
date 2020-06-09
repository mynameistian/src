package model

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

// InitRedis 初始化数据库连接池
func InitRedis(maxIdle int, maxActive int, idleTimeout time.Duration, addr string) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数， 0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

// GetPoolConn 获取链接句柄
func GetPoolConn() (conn redis.Conn, err error) {
	conn = pool.Get()
	return
}
