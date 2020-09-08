package job

import (
	"testing"
	"time"
)

type TestRunner struct{}

func (r *TestRunner) Run() {}

func TestWithInterval(t *testing.T) {

	t.Run("expect interval to be set with option", func(t *testing.T) {

		expected := 10 * time.Hour

		job := New(&TestRunner{}, WithInterval(expected))

		if job.Interval != expected {
			t.Errorf("expected interval to be: %d, got: %d", expected, job.Interval)
		}

	})

}

func TestRunImmediately(t *testing.T) {

	t.Run("expect run immediately to be set with option", func(t *testing.T) {

		job := New(&TestRunner{}, RunImmediately())

		if !job.RunImmediately {
			t.Errorf("expected run immediately to be: true, got: %t", job.RunImmediately)

		}

	})

}
