package internal

import (
	"github.com/harluo/boot"
	"github.com/harluo/di"
	"github.com/hetue/boot/internal/internal/internal"
	"github.com/hetue/boot/internal/internal/param"
)

type Booter struct {
	param *param.Starter
}

func NewBooter(param *param.Starter) *Booter {
	return &Booter{
		param: param,
	}
}

func (b *Booter) Boot(constructor any) {
	di.New().Instance().Put(constructor).Build().Apply() // 注入所有步骤

	application := boot.New()
	if "" != b.param.Name {
		application.Name(b.param.Name)
	}
	if "" != b.param.Usage {
		application.Usage(b.param.Usage)
	}
	if "" != b.param.Copyright {
		application.Copyright(b.param.Copyright)
	}
	for key, value := range b.param.Metadata {
		application.Metadata(key, value)
	}

	application.Instance().Run(internal.NewBooter) // 启动应用
}
