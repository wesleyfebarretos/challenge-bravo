package task

import (
	"fmt"

	"github.com/wesleyfebarretos/challenge-bravo/job/internal/enum"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/scheduler"
)

type CurrencyUpdater struct{}

func (j CurrencyUpdater) Start() {

	scheduler := scheduler.New()
	job := func() {
		fmt.Println("hello world")
	}

	scheduler.AddTask(enum.CurrencyUpdaterTask, "@every 5s", job)
}

func NewCurrencyUpdater() CurrencyUpdater {
	return CurrencyUpdater{}
}
