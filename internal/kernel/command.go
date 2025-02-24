package kernel

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/hetue/core/internal/config"
)

type Command struct {
	*gex.Command

	_ gox.Pointerized
}

func newCommand(ctx *Context, logging *config.Logging, runtime *config.Runtime) *Command {
	builder := gex.New("").Context(ctx)
	if nil == runtime.Pwe || *runtime.Pwe {
		builder.Pwe()
	}
	if logging.Verbose {
		builder.Echo()
	}

	return &Command{
		Command: builder,
	}
}
