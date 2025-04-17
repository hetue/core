package internal

import (
	"github.com/harluo/boot"
	"github.com/hetue/boot/internal/internal/internal/internal"
	"github.com/hetue/boot/internal/internal/internal/internal/command"
)

type Starter struct {
	boot.Lifecycle

	run *command.Run
}

func NewStarter(bootstrap internal.Starter) boot.Starter {
	return &Starter{
		run: bootstrap.Run,
	}
}

func (s *Starter) Startup(application *boot.Application) error {
	return application.Add(s.run)
}
