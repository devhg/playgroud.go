package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":5200")
	if err != nil {
		panic(err)
	}
	go send(conn)
	go recv(conn)

	time.Sleep(30 * time.Second)
}

// 用户发送消息
func send(conn net.Conn) {
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			panic(err)
		}
		fmt.Println("发送：", "hello")
		time.Sleep(2 * time.Second)
	}
}

// 用户接收消息
func recv(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Println("收到：", string(data[:n]))
	}
}
