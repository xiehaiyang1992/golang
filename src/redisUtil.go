package main

import (
	"github.com/go-redis/redis"
)

var redisUtil *redis.Client

func redisInit() {
	redisUtil = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     3,
		MinIdleConns: 2,
	})
}
