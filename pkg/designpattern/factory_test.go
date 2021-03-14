package designpattern

import (
	"fmt"
	"testing"
)

// https://juejin.cn/post/6859015515344633863#heading-7
// 1.代码可读性好   2.与使用者代码解耦
// 两种方式：
// （1）提供一个工厂对象，通过调用工厂对象的工厂方法来创建产品对象；
type Type uint8

const (
	Start = iota
	End
)

// 事件接口
type Event interface {
	Type() Type
	Content() string
}

// 开始事件
type StartEvent struct {
	content string
}

func (e StartEvent) Type() Type {
	return Start
}

func (e StartEvent) Content() string {
	panic("implement me")
}

// 结束事件
type EndEvent struct {
	content string
}

func (e EndEvent) Type() Type {
	return End
}

func (e EndEvent) Content() string {
	panic("implement me")
}

// 工厂
type Factory struct{}

func (f *Factory) Create(eType Type) Event {
	switch eType {
	case Start:
		return &StartEvent{content: "this is a start event"}
	case End:
		return &EndEvent{content: "this is a end event"}
	}
	return nil
}

func TestEventFactory(t *testing.T) {
	f := &Factory{}
	start := f.Create(Start)
	end := f.Create(End)
	fmt.Println(start)
	fmt.Println(end)
}

//（2）将工厂方法集成到产品对象中（C++/Java中对象的static方法，Go中同一package下的函数）
func OfStart() Event {
	return &StartEvent{content: "a"}
}

func OfEnd() Event {
	return &EndEvent{content: "b"}
}

func TestEventFactory2(t *testing.T) {
	if start := OfStart(); start.Type() != Start {
		t.Error("err")
	}
	if end := OfEnd(); end.Type() != End {
		t.Error("err")
	}
}
