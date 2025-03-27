package command

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/log"
	"github.com/hetue/core/internal/config"
	"github.com/hetue/core/internal/internal/internal/internal/command/internal/get"
	"github.com/hetue/core/internal/kernel"
	"github.com/pangum/pangu"
)

type Run struct {
	*pangu.Command

	steps  []kernel.Step
	retry  *config.Retry
	logger log.Logger
}

func newRun(run get.Run) *Run {
	return &Run{
		Command: pangu.NewCommand("run").Aliases("r").Usage("运行").Build(),

		steps:  run.Steps,
		logger: run.Logger,
	}
}

func (r *Run) Run(ctx context.Context) (err error) {
	now := time.Now()
	fields := gox.Fields[any]{
		field.New("steps", r.steps),
	}
	r.logger.Info("插件执行开始", fields...)

	// 设置整体超时时间
	timeout, cancel := context.WithTimeout(ctx, r.retry.Timeout)
	defer cancel()

	// 执行插件
	waiter := new(sync.WaitGroup)
	for _, step := range r.steps {
		err = r.run(&timeout, step, waiter)
		if nil != err && !step.Retryable() {
			break
		}
	}
	waiter.Wait()

	elapsed := time.Now().Sub(now).Truncate(time.Second)
	fields = fields.Add(field.New("elapsed", elapsed.String()))
	if nil != err {
		r.logger.Error("插件执行出错", field.Error(err))
	} else {
		r.logger.Info("插件执行成功")
	}

	return
}

func (r *Run) run(ctx *context.Context, step kernel.Step, waiter *sync.WaitGroup) (err error) {
	if step.Asyncable() {
		err = r.async(ctx, step, waiter)
	} else {
		err = r.sync(ctx, step)
	}

	return
}

func (r *Run) sync(ctx *context.Context, step kernel.Step) error {
	return r.exec(ctx, step)
}

func (r *Run) async(ctx *context.Context, step kernel.Step, wg *sync.WaitGroup) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = r.exec(ctx, step); nil != err {
			panic(err)
		}
	}()

	return
}

func (r *Run) exec(ctx *context.Context, step kernel.Step) (err error) {
	if !step.Runnable() {
		return
	}

	counts := r.retry.Counts
	retry := step.Retryable()
	fields := gox.Fields[any]{
		field.New("name", step.Name()),
		field.New("async", step.Asyncable()),
		field.New("retry", retry),
		field.New("counts", counts),
	}

	r.logger.Info("步骤执行开始", fields...)
	for count := 0; count < counts; count++ {
		if err = step.Run(ctx); (nil == err) || (0 == count && !retry) {
			break
		}

		backoff := rand.New().Duration().Between(time.Second, r.retry.Backoff).Build().Generate().Truncate(time.Second)
		r.logger.Info(fmt.Sprintf("步骤第%d次执行遇到错误", count+1), fields.Add(field.Error(err))...)
		r.logger.Info(fmt.Sprintf("休眠%s，继续执行步骤", backoff), fields...)
		time.Sleep(backoff)
		r.logger.Info(fmt.Sprintf("步骤重试第%d次执行", count+2), fields...)

		if count != counts-1 {
			err = nil
		}
	}

	switch {
	case nil != err && retry:
		r.logger.Error("步骤执行尝试所有重试后出错", fields.Add(field.Error(err))...)
	case nil != err:
		r.logger.Error("步骤执行出错", fields.Add(field.Error(err))...)
	default:
		r.logger.Info("步骤执行成功", fields...)
	}

	return
}
