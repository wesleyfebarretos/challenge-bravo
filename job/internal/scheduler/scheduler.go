package scheduler

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	once sync.Once
)

type Task struct {
	ID      cron.EntryID `json:"id"`
	Name    string       `json:"name"`
	PrevRun time.Time    `json:"prev_run"`
	NextRun time.Time    `json:"next_run"`
}

type AvailableTask struct {
	ID      cron.EntryID `json:"-"`
	Name    string       `json:"name"`
	Removed bool         `json:"removed"`
	LastRun *time.Time   `json:"last_run"`
}

type Scheduler struct {
	cron           *cron.Cron
	tasks          []Task
	availableTasks []AvailableTask
}

var scheduler *Scheduler

func New() *Scheduler {
	once.Do(func() {
		scheduler = &Scheduler{
			cron:           cron.New(),
			tasks:          []Task{},
			availableTasks: []AvailableTask{},
		}
	})

	return scheduler
}
