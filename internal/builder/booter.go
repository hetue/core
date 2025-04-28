package builder

import (
	"github.com/hetue/boot/internal/internal"
	"github.com/hetue/boot/internal/internal/param"
)

type Booter struct {
	param *param.Starter
}

func NewBooter() *Booter {
	return &Booter{
		param: param.NewStarter(),
	}
}

func (b *Booter) Build() *internal.Booter {
	return internal.NewBooter(b.param)
}
