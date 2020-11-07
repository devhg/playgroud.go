package gredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Subscribe(channel string) {
	conn := RedisConn.Get()
	defer conn.Close()

	pubSubConn := redis.PubSubConn{Conn: conn}
	pubSubConn.Subscribe(channel) // 订阅频道

	for {
		switch v := pubSubConn.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}
