package scheduler

import (
	"slices"

	"github.com/robfig/cron/v3"
)

func (s *Scheduler) RemoveTask(taskID cron.EntryID) {
	var index int

	for i, v := range s.tasks {
		if v.ID == taskID {
			index = i
			break
		}
	}

	s.tasks = slices.Delete(s.tasks, index, index+1)

	s.cron.Remove(taskID)
}
