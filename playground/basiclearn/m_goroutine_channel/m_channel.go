package m_goroutine_channel

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c // 这里c 是闭包；实际中，chan可以作为参数和返回值
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond) // 防止最后一个数据打印不出来哦
}

func workers(id int, c chan int) {
	for {
		// 解决channel关闭，仍然接收收到空框的问题
		if ch, ok := <-c; ok {
			fmt.Printf("worker %d, received %c\n", id, ch)
		} else {
			break
		}
		//效果同上
		//for n := range c {
		//	fmt.Printf("worker %d, received %c\n", id, n)
		//}

	}
}

//创建多个channel 与 协程
func chanDemo1() {
	var chans [10]chan int
	// 创建worker  每个worker有自己的动态更新的工作单（channel）
	for i := 0; i < 10; i++ {
		chans[i] = make(chan int)
		go workers(i, chans[i])
	}

	// 通过工作单动态分发任务
	for i := 0; i < 10; i++ {
		chans[i] <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		chans[i] <- i + 'A'
	}

	//由于调度器调度，打印乱序
	time.Sleep(time.Millisecond) // 防止最后一个数据打印不出来哦
}

//更优雅的实现方式【只写channel】send only；  <-chan 、 chan<-
func createWorkers(id int) chan<- int {
	c := make(chan int)
	go func() { // 用到闭包
		for {
			fmt.Printf("worker %d, receiverd %c\n", id, <-c)
		}
	}()
	return c
}
func chanDemo2() {
	var chans [10]chan<- int
	// 创建worker  每个worker有自己的动态更新的工作单（channel）
	for i := 0; i < 10; i++ {
		chans[i] = createWorkers(i)
	}

	// 通过工作单动态分发任务
	for i := 0; i < 10; i++ {
		chans[i] <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		chans[i] <- i + 'A'
	}

	//由于调度器调度，打印乱序
	time.Sleep(time.Millisecond) // 防止最后一个数据打印不出来哦
}

//bufferChannel 带缓冲区的channel
func bufferChannel() {
	c := make(chan int, 3)

	go workers(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c' // 至此不会deadlock  因为3个缓冲位

	c <- 'd'

	close(c) // 在sleep时间内会收到一堆 空框
	time.Sleep(time.Millisecond)
}
