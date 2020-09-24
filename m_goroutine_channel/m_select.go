package m_goroutine_channel

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func selectDemo() {
	var c1, c2 = generator(), generator()
	for {
		// 谁先收到先打印谁
		select {
		case n := <-c1:
			fmt.Println("received from c1, ", n)
		case n := <-c2:
			fmt.Println("received from c2, ", n)
		default:
			fmt.Println("timeout")
		}
	}
}
