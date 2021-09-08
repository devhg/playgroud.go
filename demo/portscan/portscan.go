package portscan

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

/*端口扫描*/

type GoPool struct {
	MaxLimit int
	token    chan struct{}
}

type poolOption func(*GoPool)

func NewGoPool(options ...poolOption) *GoPool {
	pool := &GoPool{}
	for _, option := range options {
		option(pool)
	}
	return pool
}

func WithMaxLimit(limit int) poolOption {
	return func(pool *GoPool) {
		pool.MaxLimit = limit
		pool.token = make(chan struct{}, limit)
		for i := 0; i < limit; i++ {
			pool.token <- struct{}{}
		}
	}
}

func (gp *GoPool) Add(f func()) {
	<-gp.token
	go func() {
		f()
		gp.token <- struct{}{}
	}()
}

func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.token
	}
	close(gp.token)
}

const (
	ip  string = "45.77.70.122"
	tcp string = "tcp"
	udp string = "udp"
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

var ans []string
var mu sync.Mutex

func main() {
	// 用于先上项目问题检测
	// curl http://127.0.0.1:6060/debug/pprof/trace?seconds=20 > trace.out
	// go tool trace trace.out
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	pool := NewGoPool(WithMaxLimit(4))
	for i := 0; i < 20000; i++ {
		a := i
		pool.Add(func() {
			Scan(ip, a, tcp)
		})
	}

	pool.Wait()
	fmt.Println(ans)
}

func Scan(ip string, port int, typ string) {
	url := fmt.Sprintf("%s:%d", ip, port)
	_, err := net.Dial(typ, url)
	if err != nil {
		fmt.Println(url, "关闭")
		return
	}
	fmt.Println(url, "打开")
	mu.Lock()
	ans = append(ans, url)
	mu.Unlock()
}
