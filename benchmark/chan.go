package benchmark

import (
	"sync"
)

func ChanMain() {
	count := 0
	var wg sync.WaitGroup
	wg.Add(3)
	quit := make(chan struct{}, 1)
	a, b, c := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			select {
			case <-a:
				if count >= Max {
					close(quit)
					return
				}
				//fmt.Print("A")
				count++
				b <- struct{}{}
			case <-quit:
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-b:
				if count >= Max {
					close(quit)
					return
				}
				//fmt.Print("B")
				count++
				c <- struct{}{}
			case <-quit:
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c:
				if count >= Max {
					close(quit)
					return
				}
				//fmt.Print("C")
				count++
				a <- struct{}{}
			case <-quit:
				return
			}
		}
	}()
	a <- struct{}{}
	wg.Wait()
}
