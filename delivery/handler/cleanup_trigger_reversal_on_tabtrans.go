package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"
	"os"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func CleanUpTriggerReversalOnTabtrans() {
	usecase := usecase.NewApexSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	cleanUpSchedulerTime := os.Getenv("app.cleanup_trigger_time")
	_, er := task.Every(1).Day().At(cleanUpSchedulerTime).Do(usecase.CleanUpTriggerReversalOnTabtrans)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Scheduler INFO: Clean up trigger-reversal scheduler running at:", cleanUpSchedulerTime)

	task.StartBlocking()

}
