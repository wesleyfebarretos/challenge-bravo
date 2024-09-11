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

type Scheduler struct {
	cron  *cron.Cron
	tasks []Task
}

var scheduler *Scheduler

func New() *Scheduler {
	once.Do(func() {
		scheduler = &Scheduler{
			cron: cron.New(),
		}
	})

	return scheduler
}
