package internal

import (
	"context"

	"github.com/harluo/boot"
	"github.com/hetue/boot/internal/internal/internal/internal"
	"github.com/hetue/boot/internal/internal/internal/internal/command"
)

type Booter struct {
	run *command.Run
}

func NewBooter(booter internal.Booter) boot.Booter {
	return &Booter{
		run: booter.Run,
	}
}

func (b *Booter) Boot(_ context.Context) error {
	return nil
}

func (b *Booter) Subcommands() []boot.Command {
	return []boot.Command{
		b.run,
	}
}
