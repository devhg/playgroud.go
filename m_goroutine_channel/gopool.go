package m_goroutine_channel

/**
用channel实现并发控制
*/

type GoPool struct {
	MaxLimit  int           // 最大的并发goroutine数量
	tokenChan chan struct{} // 创建MaxLimit数量的token缓冲chan，用来阻塞创建协程
}

type GoPoolOption func(gp *GoPool)

func WithLimitGoPool(limit int) GoPoolOption {
	return func(gp *GoPool) {
		gp.MaxLimit = limit
		gp.tokenChan = make(chan struct{}, limit)

		for i := 0; i < limit; i++ {
			gp.tokenChan <- struct{}{}
		}
	}
}

func NewGoPool(op ...GoPoolOption) *GoPool {
	gp := &GoPool{}
	for _, option := range op {
		option(gp)
	}
	return gp
}

func (gp *GoPool) Submit(fn func()) {
	// 每提交一个协程请求，从tokenChan获取一个token，
	// 如果没有可用token阻塞
	token := <-gp.tokenChan
	go func() {
		fn()
		// 执行完一个fn()，归还token令牌
		gp.tokenChan <- token
	}()
}

// 等待所有令牌归还后，关闭chan
func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.tokenChan
	}
	close(gp.tokenChan)
}

func (gp *GoPool) size() int {
	return len(gp.tokenChan)
}

//func main() {
//	pool := NewGoPool(WithLimitGoPool(3))
//	defer pool.Wait()
//
//	for i := 0; i < 100; i++ {
//		pool.Submit(func() {
//			fmt.Println(i)
//		})
//	}
//}
