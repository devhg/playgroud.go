package main

import (
	"fmt"
	"github.com/devhg/LearnGo/demo/tcp/proto"
	"net"
	"time"
)

//func main() {
//	conn, err := net.Dial("tcp", ":8080")
//	if err != nil {
//		fmt.Println("connect to server err:", err)
//		return
//	}
//	defer conn.Close()
//	inputReader := bufio.NewReader(os.Stdin)
//	for {
//		input, _ := inputReader.ReadString('\n')
//		inputInfo := strings.Trim(input, "\r\n")
//		if strings.ToUpper(inputInfo) == "Q" {
//			return
//		}
//
//		fmt.Println("向server 发送：", inputInfo)
//		_, err := conn.Write([]byte(inputInfo))
//		if err != nil {
//			fmt.Println("write to server err:", err)
//			return
//		}
//
//		buf := [512]byte{}
//		n, err := conn.Read(buf[:])
//		if err != nil {
//			fmt.Println("recv failed, err:", err)
//			return
//		}
//		fmt.Println("从server 收到回复：", string(buf[:n]))
//	}
//}

func main() {
	//conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
	//	IP:   net.IPv4(0, 0, 0, 0),
	//	Port: 8080,
	//})
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("connect to server err:", err)
		return
	}
	defer conn.Close()

	//一次性发送多个消息，会出现tcp粘包现象，这样会导致多条数据粘在一起
	//主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。
	for i := 0; i < 20; i++ {
		msg := `hello, server`
		encode, err := proto.Encode(msg)
		if err != nil {
			panic(err)
		}
		conn.Write(encode)
	}
	time.Sleep(2 * time.Second)
}
