//package kafka

package main

import (
	"os/signal"
	"runtime"
	"syscall"

	"github.com/zngw/kafka"
	"github.com/zngw/log"
)

func main() {
	// 初始化日志
	err := log.Init("log.txt", nil)
	if err != nil {
		panic(err)
	}

	// 初始化生产生
	err = kafka.InitProducer("127.0.0.1:9092")
	if err != nil {
		panic(err)
	}
	defer kafka.Close()

	// 发送测试消息
	kafka.Send("Test", "This is Test Msg")
	kafka.Send("Test", "Hello Guoke")

	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}
