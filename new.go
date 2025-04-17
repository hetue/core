package boot

import (
	"github.com/hetue/boot/internal/builder"
)

func New() *builder.Starter {
	return builder.NewStarter()
}
