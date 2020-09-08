package job

import "time"

type Option func(j *Job)

func WithInterval(interval time.Duration) Option {
	return func(j *Job) {
		j.Interval = interval
	}
}

func RunImmediately() Option {
	return func(j *Job) {
		j.RunImmediately = true
	}
}
