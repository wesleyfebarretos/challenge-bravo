package scheduler

func (s *Scheduler) AddTask(name, spec string, fn func()) error {
	taskID, err := s.cron.AddFunc(spec, fn)
	if err != nil {
		return err
	}

	task := s.cron.Entry(taskID)

	s.tasks = append(s.tasks, Task{
		ID:      taskID,
		Name:    name,
		PrevRun: task.Prev,
		NextRun: task.Next,
	})

	return nil
}
