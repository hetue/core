package kernel

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox"
)

type Command struct {
	gox.Pointerized

	command *gex.Command
}

func newCommand() *Command {
	return &Command{
		command: gex.New(""),
	}
}
