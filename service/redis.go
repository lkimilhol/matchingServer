package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func getRedisConn() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("redis conn error")
	}
	return c
}

func close(c redis.Conn) {
	err := c.Close()
	if err != nil {
		fmt.Println("redis close error")
	}
}
