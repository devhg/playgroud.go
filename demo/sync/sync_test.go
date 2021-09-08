package sync

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// waitGroup
func TestRun1(t *testing.T) {
	// ...
}

// once
func TestRun2(t *testing.T) {
	foo := func() {
		fmt.Println(" i am foo")
	}

	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			once.Do(foo)
		}(i)
	}

	wg.Wait()
}

// 互斥锁Mutex 和 读写锁RWMutex
// rwm的读锁被某个读取者持有，为了防止后续写入者得不到写锁，后续发生在某个被阻塞的获取写锁 之后的所有读锁都将被阻塞
// rwm的写锁被某个写入者持有，为了防止后续读取者得不到读锁，后续发生在写锁释放之前的 获取读锁的行为都会获得读锁，即使这个过程中依然存在写锁请求
func TestRun3(t *testing.T) {
	var m sync.RWMutex
	go func() {
		m.RLock()
		fmt.Println("a")
		time.Sleep(time.Second)
		m.RUnlock()
	}()

	go func() {
		time.Sleep(time.Second * 1 / 4)
		m.Lock()
		fmt.Println("b")
		time.Sleep(time.Second)
		m.Unlock()
	}()

	go func() {
		time.Sleep(time.Second * 2 / 4)
		m.Lock()
		fmt.Println("c")
		m.Unlock()
	}()

	go func() {
		time.Sleep(time.Second * 3 / 4)
		m.RLock()
		fmt.Println("d")
		m.RUnlock()
	}()
	time.Sleep(time.Second * 4)
}

// sync.Cond 协程间通知
func TestRun4(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d)
			cond.L.Lock() // 下面操作必须在互斥锁中，防止数据竞争
			values[i] = fmt.Sprintf("%c", 'a'+i)
			cond.L.Unlock()
			cond.Broadcast()
		}(i)
	}

	checkCondtion := func() bool {
		fmt.Println(values)
		for i := 0; i < len(values); i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	// 下面操作必须在互斥锁中，被唤醒一次，调用一次checkCondition
	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondtion() {
		cond.Wait()
	}
}
