package abstractfactory

import (
	"github.com/devhg/LearnGo/pkg/designpattern/abstractfactory/pipeline"
	"testing"
)

// https://juejin.cn/post/6859015515344633863#heading-10
// 抽象工厂模式学习

func TestAbstractFactory(t *testing.T) {
	p := pipeline.Of(pipeline.DefaultConfig)
	p.Exec()
}
