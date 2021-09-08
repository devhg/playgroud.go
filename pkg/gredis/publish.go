package gredis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Publish(channel, message string) {
	conn := RedisConn.Get()
	defer conn.Close()

	s, err := redis.Int64(conn.Do("publish", channel, message))
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
}
