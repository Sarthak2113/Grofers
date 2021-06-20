package main

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

// RedisConnect connects to a default redis server at port 6379
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	HandleError(err)
	return c
}

// HandleError conveniently handles error.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
