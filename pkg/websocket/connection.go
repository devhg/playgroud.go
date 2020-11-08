package main

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex // 锁 保证close线程安全
	isClosed bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte),
	}

	// 开启读入协程
	go conn.readLoop()
	// 开启写出协程
	go conn.writeLoop()
	return
}

// 对外读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 对外写消息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 关闭连接
func (conn *Connection) Close() {
	// 线程安全的close
	conn.wsConn.Close()

	conn.mutex.Lock()
	if !conn.isClosed {
		// 保证closeChan只关闭一次
		close(conn.closeChan) // close chan 可以实现多地方阻塞关闭
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现读消息
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		// 阻塞在这里等待 inChan 有空闲的位置
		select {
		case conn.inChan <- data:
			// inChan没有阻塞的时候
		case <-conn.closeChan:
			// closeChan 关闭的时候
			goto ERR
		}

	}
ERR:
	conn.Close()
}

// 内部实现发送消息
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
			// outChan没有阻塞的时候
		case <-conn.closeChan:
			// closeChan 关闭的时候
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
