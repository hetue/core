package core

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/hetue/core/internal/config"
	"github.com/hetue/core/internal/kernel"
)

type Command struct {
	context *kernel.Context
	logging *config.Logging
	runtime *config.Runtime

	_ gox.Pointerized
}

func newCommand(ctx *kernel.Context, logging *config.Logging, runtime *config.Runtime) *Command {
	return &Command{
		context: ctx,
		logging: logging,
		runtime: runtime,
	}
}

func (c *Command) New(name string) (command *gex.Command) {
	command = gex.New(name).Context(c.context)
	if *c.runtime.Pwe {
		command.Pwe()
	}
	if c.logging.Verbose {
		command.Echo()
	}

	return
}
