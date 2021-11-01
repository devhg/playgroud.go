package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/devhg/LearnGo/demo/tcp/proto"
)

func main() {
	// listen, err := net.ListenTCP("tcp", &net.TCPAddr{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Port: 8080,
	// })
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listened failed, err:", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
		}
		// 起一个协程处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// var buf [128]byte
		// n, err := reader.Read(buf[:])
		recvStr, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("read from conn err:", err)
			break
		}
		// recvStr := string(buf[:n])
		fmt.Println("从client收到消息", recvStr)
		// conn.Write([]byte(recvStr))
	}
}
