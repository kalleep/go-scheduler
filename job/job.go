package job

import (
	"time"
)

const (
	defaultInterval       = 1 * time.Second
	defaultRunImmediately = false
)

type Runner interface {
	Run()
}

func New(runner Runner, options ...Option) Job {

	job := &Job{
		runner:         runner,
		Interval:       defaultInterval,
		RunImmediately: defaultRunImmediately,
	}

	for _, option := range options {
		option(job)
	}

	return *job

}

type Job struct {
	runner Runner
	RunImmediately bool
	Interval       time.Duration
}

func (j Job) Run() {
	j.runner.Run()
}
