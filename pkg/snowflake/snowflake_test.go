package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

func TestSnowflake(t *testing.T) {
	var wg sync.WaitGroup
	worker := NewWorker(5, 5)
	ch := make(chan uint64, 10000)
	count := 10000
	defer close(ch)
	wg.Add(count)
	//并发 count个goroutine 进行 snowFlake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			id, _ := worker.NextID()
			ch <- id
		}()
	}

	wg.Wait()

	m := make(map[uint64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			fmt.Printf("repeat id %d = %d\n", id, m[id])
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i + 1

		fmt.Printf("id %d = %d\n", id, m[id])
	}
	// 成功生成 snowflake ID
	fmt.Println("All", len(m), "snowflake ID Get successed!")
}
