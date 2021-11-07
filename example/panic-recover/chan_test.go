package panic

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 礼貌的关闭chan, 不能保证 并发关闭 和 发送数据 的数据竞争问题
type mchan struct {
	C chan struct{}
	sync.Once
}

func NewMchan() *mchan {
	return &mchan{C: make(chan struct{})}
}

func (mc *mchan) safeClose() {
	mc.Do(func() {
		close(mc.C)
	})
}

// 优雅的关闭chan

// 情形一： m个接受者，一个发送者，发送者通过关闭通道 来发送结束信号
func TestRun1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	log.SetFlags(0)

	const Max = 10000
	const Receivers = 100

	wg := sync.WaitGroup{}
	wg.Add(Receivers)

	datac := make(chan int)

	go func() {
		for {
			value := rand.Intn(Max)
			if value == 0 {
				close(datac)
				return
			}
			datac <- value
		}
	}()

	for i := 0; i < Receivers; i++ {
		go func() {
			defer wg.Done()
			for value := range datac {
				log.Println(value)
			}
		}()
	}
	wg.Wait()
}

// 情形二： 一个接收者，m个发送者，此接收者 通过关闭一个额外的通道信号，来通知发送者不要继续发送数据了
func TestRun2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const Max = 10000
	const Senders = 1000

	datac := make(chan int)
	stopc := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range datac {
			if value == Max-1 {
				close(stopc) // 所有的发送者 都会受到一个零值
				return
			}
			log.Println(value)
		}
	}()

	for i := 0; i < Senders; i++ {
		go func() {
			for {
				// defer wg.Done() // 这里不要添加到waitgroup，因为他的关闭受stopc控制
				select {
				case <-stopc:
					return
				default:
				}

				select {
				case <-stopc:
					return
				case datac <- rand.Intn(Max):
				}
			}
		}()
	}
	wg.Wait()
}

// 情形三：m个接收者，n个发送者。他们任何一个协程都可以让一个 中间协调协程 帮忙发出终止信号
func TestRun3(t *testing.T) {

}
