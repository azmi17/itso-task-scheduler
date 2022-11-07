package handler

import (
	"errors"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/usecase"

	"github.com/go-co-op/gocron"
	"github.com/kpango/glg"
)

func FeeUpdateTelkomHalloTransactionOnRekpon() {
	usecase := usecase.NewRekponSchedulerUsecase()

	localTime := helper.IDNLocalTime()

	s := gocron.NewScheduler(localTime)

	_, er := s.Every(50).Second().Do(usecase.RekponUpdateFeeOnTelkomHalloTrans)
	if er != nil {
		_ = glg.Log(errors.New(er.Error()))
	}
	_ = glg.Log("Update fee telkom & hallo transaction scheduler running at: every 50 sec..")

	s.StartBlocking()

}
