package internal

import (
	"github.com/hetue/core/internal/internal/internal/internal/command"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Get

	Run *command.Run
}
