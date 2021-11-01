package safemap

import "context"

// 玩具:
// 概述：利用一个单独的goroutine管理map 实现 并发安全的map
// 缺点：性能极差

const (
	ADD = iota
	GET
	DEL
)

type SafeMap struct {
	ctx    context.Context
	cancel context.CancelFunc
	closed bool

	data map[interface{}]interface{}
	op   chan [3]interface{}
	res  chan interface{}
}

func NewSafeMap() *SafeMap {
	mp := &SafeMap{}
	mp.init()
	go mp.run()
	return mp
}

func (sm *SafeMap) init() {
	sm.ctx, sm.cancel = context.WithCancel(context.Background())
	sm.data = make(map[interface{}]interface{})
	sm.op = make(chan [3]interface{})
	sm.res = make(chan interface{})
}

func (sm *SafeMap) run() {
	for {
		select {
		case op := <-sm.op:
			switch op[0] {
			case ADD:
				sm.add(op[1], op[2])
			case GET:
				sm.get(op[1])
			case DEL:
				sm.delete(op[1])
			}
		case <-sm.ctx.Done():
			return
		}
	}
}

func (sm *SafeMap) add(k, v interface{}) {
	sm.data[k] = v
}

func (sm *SafeMap) get(k interface{}) {
	if v, ok := sm.data[k]; ok {
		sm.res <- v
		return
	}
	sm.res <- nil
}

func (sm *SafeMap) delete(k interface{}) {
	delete(sm.data, k)
}

func (sm *SafeMap) opWrap(op int, k, v interface{}) [3]interface{} {
	return [3]interface{}{op, k, v}
}

func (sm *SafeMap) close() {
	sm.cancel()
	sm.closed = true
}

// 对外接口
func (sm *SafeMap) Add(k, v interface{}) {
	if sm.closed {
		panic("Add to a closed SafeMap")
	}
	sm.op <- sm.opWrap(ADD, k, v)
}

func (sm *SafeMap) Get(k interface{}) (interface{}, bool) {
	if sm.closed {
		panic("Get from a closed SafeMap")
	}
	sm.op <- sm.opWrap(GET, k, nil)
	if ret := <-sm.res; ret != nil {
		return ret, true
	}
	return nil, false
}

func (sm *SafeMap) Delete(k interface{}) {
	if sm.closed {
		panic("Delete from a closed SafeMap")
	}
	sm.op <- sm.opWrap(DEL, k, nil)
}

func (sm *SafeMap) Close() {
	if sm.closed {
		panic("close a closed SafeMap")
	}
	sm.close()
}
