package m_goroutine_channel

import (
	"fmt"
	"os"
	"time"
)

func timeout1() {
	t := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		t <- true
	}()

	ints := make(chan int, 1)
	go func() {
		count := 0
		for {
			ints <- count
			count++
			time.Sleep(time.Duration(count) * time.Second)
		}
	}()

	for {
		select {
		case out := <-t:
			fmt.Println("超时", out)
			os.Exit(0)
		case val := <-ints:
			fmt.Println(val)
		}
	}
}

// 定时操作，推荐使用
// counter * second定时秒数。
// 超时操作，每5秒打印一次timeout。
func timeout() {
	notify := make(chan int64, 1)
	timeoutCnt := 0
	go func() {
		counter := 0
		for {
			time.Sleep(1 * time.Second)
			counter++
			notify <- time.Now().UnixNano()
			if counter > 3 {
				break
			}
		}
	}()

	for {
		select {
		case currentTime := <-notify:
			fmt.Println(currentTime)
		case <-time.After(5 * time.Second):
			fmt.Println("time out")
			timeoutCnt++
			if timeoutCnt > 3 {
				os.Exit(0)
			}
		}
	}
}
