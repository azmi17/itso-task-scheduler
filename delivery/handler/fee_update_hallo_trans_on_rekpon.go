package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func FeeUpdateHalloTransOnRekpon() {
	usecase := usecase.NewRekponSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	task := gocron.NewScheduler(localTime)

	_, er := task.Every("2m").Do(usecase.RekponFeeUpdateOnHalloTrans)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Update fee hallo transaction scheduler running at: every 2 minutes")

	task.StartBlocking()

}
