package m_io

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
)

func producer(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(100)
	data <- n
	wg.Done()
}

func consumer(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		println(err)
		f.Close()
		return
	}

	for n := range data {
		_, err := fmt.Fprintln(f, n)
		if err != nil {
			println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		println(err)
		done <- false
		return
	}
	done <- true
}

// 读取文件
func TestIO() {
	var filename = "text.txt"
	if file, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		fmt.Println(string(file))
	}
}

// 逐行读取
func TestIOByLine() {
	var filename = "text.txt"
	open, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(open)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func main() {
//	data := make(chan int)
//	done := make(chan bool)
//
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go producer(data, &wg)
//	}
//	go consumer(data, done)
//	go func() {
//		wg.Wait()
//		close(data)
//	}()
//	d := <-done
//	if d == true {
//		println("success")
//	} else {
//		println("error")
//	}
//
//}
