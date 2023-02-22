package patterns

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// 发送的信息
	interrupt chan os.Signal

	// 报告任务已完成
	complete chan error

	// 报告任务超时
	timeout <-chan time.Time

	// 函数
	tasks []func(int)
}

var ErrTimeOut = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// 接收所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当前任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.goInterrupt() {
			return ErrInterrupt
		}
		// 执行注册过的任务
		task(id)
	}
	return ErrTimeOut
}

func (r *Runner) goInterrupt() bool {
	select {
	// 当中断事件被出发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}
