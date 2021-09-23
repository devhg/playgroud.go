package abstractfactory

import (
	"testing"

	"github.com/devhg/playgroud.go/pkg/designpattern/abstractfactory/pipeline"
)

// https://juejin.cn/post/6859015515344633863#heading-10
// 抽象工厂模式学习

func TestAbstractFactory(t *testing.T) {
	p := pipeline.Of(pipeline.DefaultConfig)
	p.Exec()
}
