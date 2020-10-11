package algorithm

import (
	"errors"
	"sync"
	"time"
)

/**
雪花算法：
41位时间戳+10位机器id(5位数据中心id+5位机器id)+12位ID序列号
*/

type Worker struct {
	mu           sync.Mutex
	LastStamp    int64 // 上一次的时间戳
	WorkerID     int64 // 当前节点ID
	DataCenterID int64 // 数据中心ID
	SequenceID   int64 // 当前毫秒已经生成的ID序列号(从0 开始累加) 1毫秒内最多生成4096个ID
}

const (
	workerIDBits     = uint64(5)  // 10bit 工作机器ID中的 5bit workerID
	dataCenterIDBits = uint64(5)  // 10bit 工作机器ID中的 5bit dataCenterID
	sequenceBits     = uint64(12) // 12位序列号，表示同一机器同一毫秒可以生成 2<<12 -1个id

	maxWorkerID     = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxSequence     = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22) // timeLeft = workerIDBits + sequenceBits // 时间戳向左偏移量
	dataLeft = uint8(17) // dataLeft = dataCenterIDBits + sequenceBits
	workLeft = uint8(12) // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1589923200000) // 常量时间戳(毫秒)
)

//分布式情况下,我们应通过外部配置文件或其他方式为每台机器分配独立的id
func NewWorker(workerID, dataCenterID int64) *Worker {
	return &Worker{
		LastStamp:    0,
		WorkerID:     workerID,
		DataCenterID: dataCenterID,
		SequenceID:   0,
	}
}

func (w *Worker) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *Worker) NextID() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.nextID()
}

func (w *Worker) nextID() (uint64, error) {
	timeStamp := w.getMilliSeconds()

	if timeStamp < w.LastStamp {
		return 0, errors.New("time is moving backwards,waiting until")
	}

	if timeStamp == w.LastStamp {
		w.SequenceID = (w.SequenceID + 1) & maxSequence
		// 溢出 等待1ms
		if w.SequenceID == 0 {
			for timeStamp <= w.LastStamp {
				timeStamp = w.getMilliSeconds()
			}
		}
	} else {
		// 切换协议毫秒 置零
		w.SequenceID = 0
	}
	w.LastStamp = timeStamp

	id := (timeStamp-twepoch)<<timeLeft |
		(w.DataCenterID << dataLeft) |
		(w.WorkerID << workLeft) |
		w.SequenceID
	return uint64(id), nil
}

/*
func main() {
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
*/
