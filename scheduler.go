package scheduler

import (
	"sync"
	"time"

	"github.com/kalleep/go-scheduler/job"
)

func New(jobs ...job.Job) *Scheduler {

	return &Scheduler{
		jobs:     jobs,
		stopChan: make(chan struct{}),
		lock:     &sync.Mutex{},
		wg:       &sync.WaitGroup{},
	}
}

// Scheduler is responsible to schedule jobs and keep track of running jobs
type Scheduler struct {
	jobs []job.Job

	scheduled     bool
	scheduledJobs int

	lock     *sync.Mutex
	wg       *sync.WaitGroup
	stopChan chan struct{}
}

func (s *Scheduler) Schedule() {

	if s.scheduled {
		return
	}

	for _, j := range s.jobs {
		s.incJobs()
		go s.schedule(j)
	}

	s.scheduled = true

}

func (s *Scheduler) schedule(j job.Job) {

	ticker := time.NewTicker(j.Interval)

	defer s.decJobs()

	if j.RunImmediately {
		j.Run()
	}

	for {
		select {
		case <-ticker.C:
			j.Run()
		case <-s.stopChan:
			return
		}
	}

}

func (s *Scheduler) incJobs() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.scheduledJobs += 1
	s.wg.Add(1)
}

func (s *Scheduler) decJobs() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.scheduledJobs -= 1
	s.wg.Done()
}

func (s *Scheduler) ScheduledJobs() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.scheduledJobs
}

// Stop will tell all jobs to stop, will wait for all jobs to complete before they are stopped
func (s *Scheduler) Stop() {
	close(s.stopChan)
	s.wg.Wait()
}
