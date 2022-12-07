package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"
	"os"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func RepostingSaldoApexByScheduler() {
	usecase := usecase.NewApexSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	schedulerTime := os.Getenv("app.time_test")
	_, er := task.Every(schedulerTime + "m").Do(usecase.RepostingSaldoByScheduler)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Scheduler INFO: Reposting saldo apex scheduler running at: every", schedulerTime, "minutes")

	go RepostingSchedulerRepoObserver()

	task.StartBlocking()

}
