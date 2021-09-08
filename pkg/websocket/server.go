package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println("Recv:", string(data))
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("http listened error")
	}
}

/*
// 测试close(chan)
func main() {
	bytes := make(chan byte)
	go loop(1, bytes)
	go loop(2, bytes)

	time.Sleep(1 * time.Second)
	close(bytes)
	time.Sleep(2 * time.Second)
}

func loop(i int, close chan byte) {
	for {
		select {
		case <-close:
			fmt.Println(i, "close")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}
*/
