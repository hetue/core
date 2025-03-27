package core

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/hetue/core/internal/config"
	"github.com/hetue/core/internal/kernel"
)

type Command struct {
	context *kernel.Context
	runtime *config.Runtime

	_ gox.Pointerized
}

func newCommand(ctx *kernel.Context, runtime *config.Runtime) *Command {
	return &Command{
		context: ctx,
		runtime: runtime,
	}
}

func (c *Command) New(name string) (command *gex.Command) {
	command = gex.New(name).Context(c.context)
	if *c.runtime.Pwe {
		command.Pwe()
	}
	if c.runtime.Verbose {
		command.Echo()
	}

	return
}
