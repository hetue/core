package kernel

import (
	"context"
)

type Step interface {
	// Name 名称
	Name() string

	// Runnable 是否需要执行本步骤
	Runnable() bool

	// Retryable 是否需要重试
	Retryable() bool

	// Asyncable 是否异步执行
	Asyncable() bool

	// Run 执行逻辑
	Run(ctx *context.Context) error
}
