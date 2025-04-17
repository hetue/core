package internal

import (
	"github.com/hetue/boot/internal/internal/internal/internal"
	"github.com/hetue/boot/internal/internal/internal/internal/command"
)

type Bootstrap struct {
	boot.Lifecycle

	run *command.Run
}

func NewBootstrap(bootstrap internal.Bootstrap) core.Bootstrap {
	return &Bootstrap{
		run: bootstrap.Run,
	}
}

func (b *Bootstrap) Startup(application *core.Application) error {
	return application.Add(b.run)
}
