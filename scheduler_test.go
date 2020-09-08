package scheduler

import (
	"testing"
	"time"

	"github.com/kalleep/go-scheduler/job"
)

type TestRunner struct{}

func (r *TestRunner) Run() {}

func TestScheduler_Schedule(t *testing.T) {

	t.Run("expect jobs to be scheduled", func(t *testing.T) {

		jobs := []job.Job{
			job.New(&TestRunner{}, job.WithInterval(1*time.Second)),
			job.New(&TestRunner{}, job.RunImmediately()),
		}

		scheduler := New(jobs...)

		scheduler.Schedule()

		defer scheduler.Stop()

		if scheduler.ScheduledJobs() != 2 {
			t.Errorf("expected running jobs: 2, got: %d", scheduler.ScheduledJobs())
		}

	})

}

func TestScheduler_Stop(t *testing.T) {

	t.Run("expect jobs to be scheduled", func(t *testing.T) {

		jobs := []job.Job{
			job.New(&TestRunner{}, job.WithInterval(1*time.Second)),
			job.New(&TestRunner{}, job.RunImmediately()),
		}

		scheduler := New(jobs...)

		scheduler.Schedule()

		scheduler.Stop()

		if scheduler.ScheduledJobs() != 0 {
			t.Errorf("expected: all jobs to be stopped : got: %d", scheduler.ScheduledJobs())
		}

	})

}
