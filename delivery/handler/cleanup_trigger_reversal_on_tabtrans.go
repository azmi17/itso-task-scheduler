package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func CleanUpTriggerReversalOnTabtrans() {
	usecase := usecase.NewApexSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	_, er := task.Every("2m").Do(usecase.CleanUpTriggerReversalOnTabtrans)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Clean up trigger-reversal scheduler running at: every 2 minutes")

	task.StartBlocking()

}
