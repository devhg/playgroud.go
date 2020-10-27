package taskpool

type GoPool struct {
	MaxLimit  int
	tokenChan chan struct{}
}

type poolOption func(*GoPool)

func NewGoPool(op ...poolOption) *GoPool {
	gp := &GoPool{}
	for _, option := range op {
		option(gp)
	}
	return gp
}

func WithMaxLimit(maxLimit int) poolOption {
	return func(gp *GoPool) {
		gp.tokenChan = make(chan struct{}, maxLimit)
		gp.MaxLimit = maxLimit
		for i := 0; i < gp.MaxLimit; i++ {
			gp.tokenChan <- struct{}{}
		}
	}
}

func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.tokenChan
	}
	close(gp.tokenChan)
}

func (gp *GoPool) Submit(fn func()) {
	token := <-gp.tokenChan
	go func() {
		fn()
		gp.tokenChan <- token
	}()
}

func (gp *GoPool) Size() int {
	return len(gp.tokenChan)
}
