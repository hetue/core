package builder

import (
	"github.com/hetue/boot/internal/internal"
	"github.com/hetue/boot/internal/internal/param"
)

type Starter struct {
	param *param.Starter
}

func NewStarter() *Starter {
	return &Starter{
		param: param.NewStarter(),
	}
}

func (b *Starter) Build() *internal.Starter {
	return internal.NewStarter(b.param)
}
