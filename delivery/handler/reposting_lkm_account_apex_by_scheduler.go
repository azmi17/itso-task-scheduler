package handler

import (
	"errors"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"
	"os"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func RepostingSaldoApexByScheduler() {
	schedulerStatus := os.Getenv("app.scheduler_status")
	if schedulerStatus != entities.SCHEDULER_DISABLE {
		usecase := usecase.NewApexSchedulerUsecase()

		localTime := helper.IDNLocalTime()
		schedulerTime := os.Getenv("app.balance_reposting_time")

		task := gocron.NewScheduler(localTime)
		// _, er := task.Every(schedulerTime + "m").Do(usecase.RepostingSaldoApexByScheduler) // => debug mode
		_, er := task.Every(1).Day().At(schedulerTime).Do(usecase.RepostingSaldoApexByScheduler)
		if er != nil {
			_ = glg.Log(errors.New(er.Error()))
		}
		// _ = glg.Log("Scheduler INFO: Reposting saldo apex scheduler running at: every", schedulerTime, "minutes")
		_ = glg.Log("Scheduler INFO: Reposting saldo apex scheduler running at:", schedulerTime[0:5], "&", schedulerTime[6:11])

		// go RepostingSchedulerRepoObserver()

		task.StartBlocking()
	}
}
