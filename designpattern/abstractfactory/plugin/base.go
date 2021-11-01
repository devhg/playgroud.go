package plugin

import (
	"fmt"
	"reflect"
	"strings"
)

type Plugin interface{}

// 输入接口
type Input interface {
	Plugin
	Recv() string
}

// 过滤器接口
type Filter interface {
	Plugin
	Process(string) string
}

// 输出接口
type Output interface {
	Plugin
	Send(string)
}

type HelloInput struct{}

func (h *HelloInput) Recv() string {
	return "Hello World"
}

type UpperFilter struct{}

func (u *UpperFilter) Process(s string) string {
	return strings.ToUpper(s)
}

type ConsoleOutput struct{}

func (h *ConsoleOutput) Send(s string) {
	fmt.Println(s)
}

// 插件名称与类型的映射关系，主要用于通过反射创建filter对象
var pluginNames = make(map[string]reflect.Type)

func init() {
	pluginNames["hello"] = reflect.TypeOf(HelloInput{})
	pluginNames["filter"] = reflect.TypeOf(UpperFilter{})
	pluginNames["output"] = reflect.TypeOf(ConsoleOutput{})
}
