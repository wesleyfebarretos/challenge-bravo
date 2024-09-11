package task

import (
	"fmt"
	"time"

	"github.com/wesleyfebarretos/challenge-bravo/job/internal/scheduler"
)

type CurrencyUpdater struct{}

func (j CurrencyUpdater) Start() {

	scheduler := scheduler.New()
	job := func() {
		fmt.Println("hello world")
	}

	scheduler.AddTask("Currency Updater", "@every 5s", job)
	scheduler.AddTask("Currency Updater1", "@every 5s", job)
	scheduler.AddTask("Currency Updater2", "@every 5s", job)
	scheduler.AddTask("Currency Updater3", "@every 5s", job)

	fmt.Println(scheduler.GetAllTasks())
	scheduler.Start()

	time.Sleep(5 * time.Second)
	fmt.Println(scheduler.GetAllTasks())

	tasks := scheduler.GetAllTasks()
	for _, v := range tasks {
		fmt.Println(v.ID)
		scheduler.RemoveTask(v.ID)
	}
	fmt.Println(scheduler.GetAllTasks())
	select {}
}

func NewCurrencyUpdater() CurrencyUpdater {
	return CurrencyUpdater{}
}
