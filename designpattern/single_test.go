package designpattern

import "sync"

// “饿汉模式”，实例在系统加载的时候就已经完成了初始化
type pool struct {
	id int
}

var globalPool = &pool{
	id: 111,
}

func DefaultPool() *pool {
	return globalPool
}

// “懒汉模式”，只有等到对象被使用的时候，才会去初始化它，从而一定程度上节省了内存。
// “懒汉模式”会带来线程安全问题，可以通过普通加锁，或者更高效的双重检验锁来优化。
// 对于“懒汉模式”，Go语言有一个更优雅的实现方式，那就是利用sync.Once，
// 它有一个Do方法，其入参是一个方法，Go语言会保证仅仅只调用一次该方法。
// (也可以自己用CAS实现这个once)
var once sync.Once
var globalPool2 *pool

func DefaultPool2() *pool {
	once.Do(func() {
		globalPool2 = &pool{
			id: 1,
		}
	})
	return globalPool2
}
