package designpattern

import (
	"fmt"
	"sync"
	"testing"
)

// 建造者模式
// https://juejin.cn/post/6859015515344633863#heading-4

type message struct {
	srcIP   string
	srcPort int

	dstIP   string
	dstPort int
	data    map[interface{}]interface{}
}

type builder struct {
	once *sync.Once
	msg  *message
}

func NewBuilder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &message{},
	}
}

func (b *builder) WithSrcIP(srcIP string) *builder {
	b.msg.srcIP = srcIP
	return b
}

func (b *builder) WithSrcPort(port int) *builder {
	b.msg.srcPort = port
	return b
}

func (b *builder) WithDstIP(dstIP string) *builder {
	b.msg.dstIP = dstIP
	return b
}

func (b *builder) WithDstPort(dstPort int) *builder {
	b.msg.dstPort = dstPort
	return b
}

func (b *builder) WithDataItem(key, val interface{}) *builder {
	// 保证data的map只会创建一次
	b.once.Do(func() {
		b.msg.data = map[interface{}]interface{}{}
	})
	b.msg.data[key] = val
	return b
}

func (b *builder) Build() *message {
	return b.msg
}

// 使用建造者模式来进行对象创建，使用者不再需要知道对象具体的实现细节，代码可读性也更好。
func TestBuilder(t *testing.T) {
	msg := NewBuilder().
		WithSrcIP("0.0.0.0").
		WithSrcPort(8080).
		WithDstIP("127.0.0.1").
		WithDstPort(2333).
		WithDataItem("key1", "val1").
		WithDataItem("key2", "val2").Build()

	fmt.Println(msg)
}
