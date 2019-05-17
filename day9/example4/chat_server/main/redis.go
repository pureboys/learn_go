package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     idleConn,
		MaxActive:   maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}
