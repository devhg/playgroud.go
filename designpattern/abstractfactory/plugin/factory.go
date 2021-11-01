package plugin

import "reflect"

type Config struct {
	Name string
}

type Type uint8

const (
	InputType Type = iota
	FilterType
	OutputType
)

// 插件抽象工厂接口
type Factory interface {
	Create(conf Config) Plugin
}

// input插件工厂对象 实现了抽象工厂接口
type InputFactory struct {
}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	typ := pluginNames[conf.Name]
	return reflect.New(typ).Interface().(Plugin)
}

type FilterFactory struct {
}

func (f *FilterFactory) Create(conf Config) Plugin {
	typ := pluginNames[conf.Name]
	return reflect.New(typ).Interface().(Plugin)
}

type OutputFactory struct {
}

func (o *OutputFactory) Create(conf Config) Plugin {
	typ := pluginNames[conf.Name]
	return reflect.New(typ).Interface().(Plugin)
}
