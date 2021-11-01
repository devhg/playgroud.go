package pipeline

import "github.com/devhg/playgroud.go/designpattern/abstractfactory/plugin"

// 定义 Pipeline 消息处理结构体
type Pipeline struct {
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

func (p *Pipeline) Exec() {
	msg := p.input.Recv()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

type Config struct {
	Input  plugin.Config
	Filter plugin.Config
	Output plugin.Config
}

var DefaultConfig = Config{
	Input:  plugin.Config{Name: "hello"},
	Filter: plugin.Config{Name: "filter"},
	Output: plugin.Config{Name: "output"},
}

// 保存用于创建Plugin的工厂实例，其中map的key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[plugin.Type]plugin.Factory)

// 初始化插件工厂对象
func init() {
	pluginFactories[plugin.InputType] = &plugin.InputFactory{}
	pluginFactories[plugin.FilterType] = &plugin.FilterFactory{}
	pluginFactories[plugin.OutputType] = &plugin.OutputFactory{}
}

// 根据 plugin.Type 返回对应 plugin.Factory 类型的工厂实例
func factoryOf(typ plugin.Type) plugin.Factory {
	return pluginFactories[typ]
}

func Of(conf Config) *Pipeline {
	pipe := &Pipeline{}
	pipe.input = factoryOf(plugin.InputType).Create(conf.Input).(plugin.Input)
	pipe.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
	pipe.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
	return pipe
}
