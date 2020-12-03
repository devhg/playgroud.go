package trace

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	// go run main.go 2> trace.out
	// go tool trace trace.out
	trace.Start(os.Stderr)
	defer trace.Stop()
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(1)

	wg.Add(52)
	for i := 0; i < 26; i++ {
		go func(id int) {
			defer wg.Done()
			//runtime.Gosched()
			fmt.Printf("%c", 'a'+id)
		}(i)

		go func(id int) {
			defer wg.Done()
			runtime.Gosched()
			fmt.Printf("%c", 'A'+id)
		}(i)
	}
	wg.Wait()
}

//func main() {
//	f, err := os.Create("trace.out")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	err = trace.Start(f)
//	if err != nil {
//		panic(err)
//	}
//	defer trace.Stop()
//
//	//...
//}
