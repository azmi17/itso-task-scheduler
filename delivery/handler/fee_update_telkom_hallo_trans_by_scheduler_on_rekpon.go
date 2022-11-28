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

func FeeUpdateTelkomHalloTransOnRekpon() {
	usecase := usecase.NewRekponSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	feeSchedulerTime := os.Getenv("app.fee_scheduler_time")
	_, er := task.Every(feeSchedulerTime + "m").Do(usecase.RekponUpdateFeeOnTelkomHalloTrans)
	if er != nil {
		entities.PrintError(er.Error())
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Scheduler INFO: Update fee telkom & hallo transaction scheduler running at: every", feeSchedulerTime, "minutes")

	task.StartBlocking()

}
