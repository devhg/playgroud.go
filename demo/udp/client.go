package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
// 		IP:   net.IPv4(0, 0, 0, 0),
// 		Port: 8088,
// 	})
// 	if err != nil {
// 		fmt.Println("connect to server, err:", err)
// 		return
// 	}

// 	defer udp.Close()
// 	//for i := 0; i < 20; i++ {
// 	// udp 没有粘包
// 	_, err = udp.Write([]byte("Hello server"))
// 	if err != nil {
// 		fmt.Println("send data err:", err)
// 		return
// 	}
// 	//}

// 	var data [1024]byte
// 	n, addr, err := udp.ReadFromUDP(data[:])
// 	if err != nil {
// 		fmt.Println("recv data err:", err)
// 		return
// 	}
// 	fmt.Printf("from addr:%v recv:%v count:%v\n", addr, string(data[:n]), n)
// }
