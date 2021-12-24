package main

import (
	"context"
	"fmt"
	"time"
)

// 发现Go语言扩展包提供了一个带权重的信号量库Semaphore，
// 使用信号量我们可以实现一个"工作池"控制一定数量的goroutine并发工作。
// 因为对源码抱有好奇的态度，所以在周末仔细看了一下这个库并进行了解析，在这里记录一下。
func main() {
	s := semaphore.NewWeighted(3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	for i := 0; i < 3; i++ {
		if i != 0 {
			go func(num int) {
				if err := s.Acquire(ctx, 3); err != nil {
					fmt.Printf("goroutine： %d, err is %s\n", num, err.Error())
					return
				}
				time.Sleep(2 * time.Second)
				fmt.Printf("goroutine： %d run over\n", num)
				s.Release(3)

			}(i)
		} else {
			go func(num int) {
				ct, cancel := context.WithTimeout(context.Background(), time.Second*3)
				defer cancel()
				if err := s.Acquire(ct, 3); err != nil {
					fmt.Printf("goroutine： %d, err is %s\n", num, err.Error())
					return
				}
				time.Sleep(3 * time.Second)
				fmt.Printf("goroutine： %d run over\n", num)
				s.Release(3)
			}(i)
		}

	}
	time.Sleep(10 * time.Second)
}
