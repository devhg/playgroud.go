package panic_test

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

//panic recover demo

/****************************************/
//自动重启因为panic而退出的进程
func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			log.Printf("last dead, reload one: %v\n", name)
			go NeverExit(name, f)
		}
	}()
	f()
}

func TestRun(t *testing.T) {
	do := func() {
		for {
			time.Sleep(time.Second)
			if time.Now().UnixNano()&0x1 == 0 {
				panic("unexpected error")
			}
		}
	}

	go NeverExit("go1", do)
	go NeverExit("go2", do)
	select {} // 永远阻塞
}

/****************************************/
//一旦panic发生，内嵌函数将跳转到defer处，可以实现跨函数跳转，但不推荐这样做
func TestRun1(t *testing.T) {
	n := func() (result int) {
		defer func() {
			if v := recover(); v != nil {
				if i, ok := v.(int); ok {
					result = i
				}
			}
		}()
		func() {
			func() {
				func() {
					func() {
						panic(123)
					}()
					fmt.Println("innner")
				}()
			}()
		}()
		return 0
	}()
	fmt.Println(n)
}

/****************************************/
//最快响应
func longReq(r chan<- int) {
	time.Sleep(time.Second * 3)
	r <- rand.Intn(100)
}

func TestRun2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a, b := make(chan int), make(chan int)

	//base mode
	startTime := time.Now()
	go longReq(a)
	go longReq(b)
	fmt.Println(<-a, <-b)
	fmt.Println("base mode:", time.Since(startTime))

	//fast mode
	ch := make(chan int, 5)
	startTime = time.Now()
	for i := 0; i < cap(ch); i++ {
		go longReq(ch)
	}
	fmt.Println(<-ch, <-ch, len(ch))
	fmt.Println("super mode:", time.Since(startTime))
}

/****************************************/
//		channel通道实现通知和定时器
/****************************************/
func one2one() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()
	// do other concurrent

	<-done
	fmt.Println(values[0], values[len(values)-1])
}

func one2more() {
	worker := func(id int, ready <-chan struct{}, done chan<- struct{}) {
		<-ready
		log.Printf("worker#, %d start work\n", id)
		time.Sleep(time.Second * 2)
		log.Printf("worker#, %d work ok\n", id)
		done <- struct{}{}
	}

	ready, done := make(chan struct{}), make(chan struct{})
	go worker(1, ready, done)
	go worker(2, ready, done)
	go worker(3, ready, done)

	// 单对多通知开始
	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}
	// 等待多对单通知
	<-done
	<-done
	<-done
	//可以用sync.WaitGroup优雅的实现
}

//通过通道的关闭实现群发通知
//从一个已经关闭的通道可以接收到无穷个值。实际上这一个特性被广泛用于标准库，如context
func one2moreByclose() {
	worker := func(id int, ready <-chan struct{}, done chan<- struct{}) {
		<-ready
		log.Printf("worker#, %d start work\n", id)
		time.Sleep(time.Second * 2)
		log.Printf("worker#, %d work ok\n", id)
		done <- struct{}{}
	}

	ready, done := make(chan struct{}), make(chan struct{})
	go worker(1, ready, done)
	go worker(2, ready, done)
	go worker(3, ready, done)

	// 单对多通知开始
	time.Sleep(time.Second * 3)
	close(ready)
	// 等待多对单通知
	<-done
	<-done
	<-done
	//可以用sync.WaitGroup优雅的实现
}

func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func TestRun3(t *testing.T) {
	//one2one()
	//one2more()
	//one2moreByclose()
	<-AfterDuration(time.Second * 2)
	fmt.Println("hello")
	//select {}
}

/****************************************/
//		channel通道实现互斥锁
func myMutex() {
	mutex := make(chan struct{}, 1)
	counter := 0
	increase := func() {
		mutex <- struct{}{}
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}
	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter) // 2000
}

//将通道用作技术信号量
//计数信号量常常被用作限制最大并发数
func countSemaphore() {
	rand.Seed(time.Now().UnixNano())
	token := make(chan struct{}, 10)

	worker := func(id int) {
		log.Println("id#", id, "开始工作")
		<-token
		time.Sleep(time.Second * 2)
		log.Println("id#", id, "工作完成")
		token <- struct{}{}
	}

	for i := 0; i < cap(token); i++ {
		token <- struct{}{}
	}

	for workerId := 1; ; workerId++ {
		time.Sleep(time.Second)
		go worker(workerId)
	}
}
func TestRun4(t *testing.T) {
	//myMutex()
	countSemaphore()
}

/****************************************/
//使用通道传输数据

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		//process
		for request := range requests {
			if request == nil {
				n++
			} else {
				request <- n
			}
		}
	}()
	return requests // 返回一个chan类型的数据缓冲区
}(0)

func TestRun5(t *testing.T) {
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done) // add
	go increase1000(done)
	<-done
	<-done // 2 goroutine add complete
	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request)
}

/****************************************/
//超时机制
func TestRun6(t *testing.T) {
	c := make(chan int)
	select {
	case data := <-c:
		fmt.Println(data)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

/****************************************/
//脉搏器
func Tick(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func TestRun7(t *testing.T) {
	//start := time.Now()
	//for range Tick(time.Second * 3) {
	//	fmt.Println(time.Since(start))
	//}
	start := time.Now()
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println(time.Since(start))
		}
	}

}
