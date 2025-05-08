package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/boot/internal/internal/internal/internal/command"
	"github.com/hetue/boot/internal/internal/internal/internal/finder"
)

type Booter struct {
	di.Get

	Run    *command.Run
	Finder finder.Handler `optional:"true"`
}
