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
		Runner:         runner,
		Interval:       defaultInterval,
		RunImmediately: defaultRunImmediately,
	}

	for _, option := range options {
		option(job)
	}

	return *job

}

type Job struct {
	Runner         Runner
	RunImmediately bool
	Interval       time.Duration
}
