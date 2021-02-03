package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

/*-------------------INPUT-------------------*/
//var reader = bufio.new(os.Stdin)
//var writer = bufio.NewWriter(os.Stdout)

//var writer = bufio.NewWriter(os.Stdout)
//var reader = bufio.NewScanner(os.Stdin)

func main() {
	parse, err := time.Parse("2006-01-02 15:04:05", "2021-02-03 21:12:40")
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithDeadline(context.Background(), parse)
	ok := make(chan int)
	begin := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done(): // 定时开始
				fmt.Println(123)
				ok <- 1
			case <-begin: // 条件开始
			}
		}
	}()
	<-ok
	ctx1, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second): // 定时退出
				cancel()
			case <-ctx1.Done(): // 条件对出
				cancel()
			default:
				time.Sleep(time.Second)
				fmt.Println("heat")
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
