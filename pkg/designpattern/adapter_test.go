package designpattern

import "fmt"

//适配器模式是最常用的结构型模式之一，它让原本因为接口不匹配而无法一起工作的两个对象能够一起工作。
//在现实生活中，适配器模式也是处处可见，比如电源插头转换器，可以让英式的插头工作在中式的插座上。

//适配器模式所做的就是将一个接口Adaptee，通过适配器Adapter转换成Client所期望的另一个接口Target来使用，
//实现原理也很简单，就是Adapter通过实现Target接口，并在对应的方法中调用Adaptee的接口实现。
//一个典型的应用场景是，系统中一个老的接口已经过时即将废弃，但因为历史包袱没法立即将老接口全部替换为新接口，
//这时可以新增一个适配器，将老的接口适配成新的接口来使用。

//适配器模式很好的践行了面向对象设计原则里的开闭原则（open/closed principle），
//新增一个接口时也无需修改老接口，只需多加一个适配层即可。

type Records struct {
	Items []string
}

type Consumer interface {
	Poll() Records
}

type KafkaInput struct {
	status   string
	consumer Consumer
}

func (k *KafkaInput) Receive() {
	records := k.consumer.Poll()
	fmt.Println(records)
}
