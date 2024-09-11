package main

import "github.com/wesleyfebarretos/challenge-bravo/job/internal/task"

func main() {
	task := task.NewCurrencyUpdater()
	task.Start()
}
