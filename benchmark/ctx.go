package benchmark

import (
	"context"
)

func CtxMain() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	a, b, c := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go func() {
		for {
			select {
			case <-a:
				//fmt.Print("A")
				done <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-b:
				//fmt.Print("B")
				done <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-c:
				//fmt.Print("C")
				done <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()
	for i := 0; i < Max; i += 3 {
		a <- struct{}{}
		<-done
		b <- struct{}{}
		<-done
		c <- struct{}{}
		<-done
	}
	cancel()
}
