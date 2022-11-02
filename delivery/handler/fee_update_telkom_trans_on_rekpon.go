package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func FeeUpdateTelkomTransOnRekpon() {
	usecase := usecase.NewRekponSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	_, er := task.Every("2m").Do(usecase.RekponFeeUpdateOnTelkomTrans)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Update fee telkom transaction scheduler running at: every 2 minutes")

	task.StartBlocking()

}
