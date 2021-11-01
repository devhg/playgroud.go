package main

// func main() {
// 	total, sum := 0, 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i
// 		go func() {
// 			total += i
// 		}()
// 	}
// 	fmt.Printf("total:%d sum %d", total, sum)
// }

// 加锁解决
// func main() {
// 	var mu sync.Mutex
// 	total, sum := 0, 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i
// 		go func(n int) {
// 			mu.Lock()
// 			total += n
// 			mu.Unlock()
// 		}(i)
// 	}
// 	time.Sleep(2 * time.Second)
// }

// 最优雅的方式
// 用于先上项目问题检测
// curl	http://127.0.0.1:6060/debug/pprof/trace?seconds=20 > trace.out
// go tool trace trace.out

// func main() {
// 	//这将使您的程序以二进制格式在文件trace.out中写入事件数据
// 	//然后可以运行go tool trace trace.out。
// 	//这将解析跟踪文件，并使用可视化程序打开浏览器。
// 	//该命令还将启动服务器，并使用跟踪数据来响应可视化操作。
// 	//在浏览器中加载初始页面后，单击“View trace”。
// 	//这将加载跟踪查看器，如上面嵌入的那样。
// 	f, err := os.Create("trace.out")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	err = trace.Start(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer trace.Stop()

// 	var total int64
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			atomic.AddInt64(&total, int64(n))
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(total)
// }
