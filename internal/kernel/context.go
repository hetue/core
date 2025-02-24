package kernel

import (
	"context"

	"github.com/goexl/gox"
)

type Context struct {
	context.Context

	_ gox.Pointerized
}

func newContext() *Context {
	return &Context{
		Context: context.Background(),
	}
}
