# Gin-Gonic Redis Pool Middleware

Simple redis pool middleware for Gin-Gonic framework.

This middleware set on each request a redis connection (new or pooled), and can be accessed with context `Get("Redis")` or `MustGet("Redis")` method.

```
package main

import (
		ggRedisPool "github.com/valenciaj/gg-redis-pool"
)

// Create server
srv := gin.Default()

dbConnStr =: "redis://@127.0.0.1:6379/0"

// Set middlewares
srv.Use(
	gin.Recovery(),
	ggRedisPool.Redis(dbConnStr),
)

// ... set other stuff ...

// Set routes
srv.GET("/", func(ctx *gin.Context) {

	// With Get method
	redisI, err := ctx.Get("Redis")
	if err != nil {
		panic(err)
	}
	// Unbox redis connection
	redis := redisI.(*DB)

	// With MustGet method
	redis := ctx.MustGet("Redis").(*DB)

	// ... do some stuff with redis connection ...

	// Response data
	ctx.HTML(200, "welcome/index.tmpl", gin.H{})
})

```
