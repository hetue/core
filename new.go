package boot

import (
	"github.com/hetue/boot/internal/builder"
)

func New() *builder.Booter {
	return builder.NewBooter()
}
