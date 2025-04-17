package builder

import (
	"github.com/hetue/boot/internal/internal"
	"github.com/hetue/boot/internal/internal/param"
)

type Bootstrap struct {
	param *param.Starter
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		param: param.NewStarter(),
	}
}

func (b *Bootstrap) Build() *internal.Starter {
	return internal.NewStarter(b.param)
}
