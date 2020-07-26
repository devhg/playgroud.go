package m_goroutine

import (
	"fmt"
	"time"
)

//channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放。
//close可以用来通知channel接收者不会再收到数据。所以即使channel中有数据也可以close而不会导致接收者收不到残留的数据。
//有些场景需要关闭通道，例如range遍历通道，如不关闭range遍历会出现死锁。
//有缓存的channel 关闭可以读取缓存内容  无缓存无法进行读取
var mchan = make(chan int, 10)
var closeChan = make(chan byte, 1)

func Loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func Send() {
	mchan <- 1
	time.Sleep(time.Second * 1)
	mchan <- 2
	time.Sleep(time.Second * 1)
	mchan <- 3
	time.Sleep(time.Second * 3)
	closeChan <- 1
}

func Receive() {
	for {
		select {
		case num := <-mchan:
			fmt.Println("num: ", num)
		case <-closeChan:
			fmt.Println("timeout...")
		}
	}
}
