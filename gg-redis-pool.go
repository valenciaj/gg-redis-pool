package gg-redis-pool

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func Redis(connection string) gin.HandlerFunc {

	pool := &redis.Pool{
		Dial:        func() (redis.Conn, error) { return redis.DialURL(connection) },
		MaxIdle:     50,
		MaxActive:   250,
		IdleTimeout: time.Minute * 5,
	}

	return func(ctx *gin.Context) {

		ctx.Set("Redis", pool.Get())
		ctx.Next()
	}
}
