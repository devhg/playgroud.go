package benchmark

import (
	"sync"
	"sync/atomic"
)

var count int32

const Max = 100

func AtomicMain() {
	count = 0
	var wg sync.WaitGroup
	wg.Add(3)
	go cat(&wg)
	go dog(&wg)
	go fish(&wg)
	wg.Wait()
}

func cat(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if atomic.LoadInt32(&count) >= Max {
			break
		}
		if atomic.LoadInt32(&count)%3 == 0 {
			//fmt.Print("A")
			atomic.AddInt32(&count, 1)
		}
	}
}

func dog(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if atomic.LoadInt32(&count) >= Max {
			break
		}
		if atomic.LoadInt32(&count)%3 == 1 {
			//fmt.Print("B")
			atomic.AddInt32(&count, 1)
		}
	}
}

func fish(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if atomic.LoadInt32(&count) >= Max {
			break
		}
		if atomic.LoadInt32(&count)%3 == 2 {
			//fmt.Print("C")
			atomic.AddInt32(&count, 1)
		}
	}
}
