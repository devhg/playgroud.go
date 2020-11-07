package gredis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup(pwd string) error {
	RedisConn = &redis.Pool{
		MaxIdle:     1000,
		MaxActive:   10,
		IdleTimeout: 6000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	reply, err := redis.String(conn.Do("SET", key, value))
	conn.Do("EXPIRE", key, time)

	return reply, err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

// Hash store

//利用redis库自带的Args 和 AddFlat对结构体进行转换。然后以hash类型存储。
//该方式实现简单，但存在最大的问题是不支持数组结构（如：结构体中内嵌结构体、数组等）。
func DoHashStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.String(conn.Do("hmset", redis.Args{key}.AddFlat(src)...))
}

func DoHashGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		return err
	}
	return redis.ScanStruct(value, dest)
}

// Gob Encoding
func DoGobStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(src)

	return redis.String(conn.Do("set", key, buffer.Bytes()))
}

func DoGobGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	reBytes, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return err
	}

	reader := bytes.NewReader(reBytes)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(dest)
}

// JSON Encoding
func DoJsonStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	datas, err := json.Marshal(src)
	if err != nil {
		return "", err
	}
	return redis.String(conn.Do("set", key, datas))
}

func DoJsonGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	datas, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return err
	}
	return json.Unmarshal(datas, dest)
}

func Pipelining() {
	conn := RedisConn.Get()
	defer conn.Close()

	//发送命令至缓冲区
	conn.Send("HSET", "student", "name", "wd", "age", "22")
	conn.Send("HSET", "student", "Score", "100")
	conn.Send("HGET", "student", "age")

	conn.Flush() //清空缓冲区，将命令一次性发送至服务器

	//依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
	if err != nil {
		panic(err)
	}
}

func transaction() {
	conn := RedisConn.Get()
	defer conn.Close()

	//MULTI：开启事务
	//EXEC：执行事务
	//DISCARD：取消事务
	//WATCH：监视事务中的键变化，一旦有改变则取消事务。

	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
	if err != nil {
		panic(err)
	}
}
