package m_goroutine_channel

import (
	"fmt"
	"sync"
	"time"
)

//线程安全的int
type atomicInt struct {
	Value int
	lock  sync.Mutex
}

func (a *atomicInt) addIncrement() {
	fmt.Println("increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.Value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.Value
}

func mutexDemo() {
	var a atomicInt
	a.addIncrement()

	go func() {
		a.addIncrement()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
