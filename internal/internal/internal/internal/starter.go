package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/boot/internal/internal/internal/internal/command"
)

type Starter struct {
	di.Get

	Run *command.Run
}
