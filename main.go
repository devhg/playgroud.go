package main

import (
	"fmt"
	"github.com/QXQZX/LearnGo/m_goroutine_channel"
	"time"
)

type Base struct {
}

type Child struct {
	Base
}

func main() {
	pool := m_goroutine_channel.NewGoPool(m_goroutine_channel.WithLimitGoPool(3))
	defer pool.Wait()

	for i := 0; i < 100; i++ {
		pool.Submit(func() {
			fmt.Println(i)
		})
		time.Sleep(1 * time.Second)
	}
}
