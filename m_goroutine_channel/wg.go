package m_goroutine_channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	id     int
	random int
}

type result struct {
	job job
	sum int
}

var jobs = make(chan job, 10)
var results = make(chan result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func allocate(num int) {
	// num=100 生成100个job  信道cap=10 信道装满阻塞
	for i := 0; i < num; i++ {
		random := rand.Intn(999)
		job := job{i, random}
		jobs <- job
	}
	close(jobs)
}
func res(done chan bool) {
	for r := range results {
		fmt.Printf("job id is %d,  sum is %d\n", r.job.id, r.sum)
	}
	done <- true
}

func worker(wg *sync.WaitGroup) {
	for j := range jobs {
		op := result{j, digits(j.random)}
		results <- op
	}
	wg.Done()
}

//函数创建了一个 Go 协程的工作池。
func createWorkBool(num int) {
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func TestWg() {
	startTime := time.Now()
	jobsNum := 100
	done := make(chan bool)

	go allocate(jobsNum)  // 生成一定数量的job
	go createWorkBool(20) //创建工作池 有多个worker工作
	go res(done)

	<-done
	endTime := time.Now()
	fmt.Println("total time taken", endTime.Sub(startTime).Seconds(), "seconds")
}
