package designpattern

import (
	"testing"
)

//原型模式主要解决对象复制的问题，它的核心就是clone()方法，返回Prototype对象的复制品。
//在程序设计过程中，往往会遇到有一些场景需要大量相同的对象，如果不使用原型模式，
//那么我们可能会这样进行对象的创建：新创建一个相同对象的实例，然后遍历原始对象的所有成员变量，
//并将成员变量值复制到新对象中。这种方法的缺点很明显，那就是使用者必须知道对象的实现细节，导致代码之间的耦合。
//另外，对象很有可能存在除了对象本身以外不可见的变量，这种情况下该方法就行不通.

//对于这种情况，更好的方法就是使用原型模式，将复制逻辑委托给对象本身，这样，上述两个问题也都迎刃而解了。
type ProtoType interface {
	clone() ProtoType
}

func (m *message) clone() ProtoType {
	msg := *m
	return &msg
}

func TestProtoType(t *testing.T) {
	msg := NewBuilder().
		WithSrcIP("0.0.0.0").
		WithSrcPort(8080).
		WithDstIP("127.0.0.1").
		WithDstPort(2333).
		WithDataItem("key1", "val1").
		WithDataItem("key2", "val2").Build()
	newMsg := msg.clone().(*message)
	if newMsg.srcIP != msg.srcIP {
		t.Errorf("clone message error")
	}
}
