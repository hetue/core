package builder

import (
	"github.com/hetue/core/internal/internal"
	"github.com/hetue/core/internal/internal/param"
)

type Bootstrap struct {
	param *param.Bootstrap
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		param: param.NewBootstrap(),
	}
}

func (b *Bootstrap) Build() *internal.Bootstrap {
	return internal.NewBootstrap(b.param)
}
