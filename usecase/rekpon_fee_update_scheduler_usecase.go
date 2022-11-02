package usecase

import (
	"fmt"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/repository/rekponrepo"
	"time"

	"github.com/kpango/glg"
)

type RekponSchedulerUsecase interface {
	RekponFeeUpdateOnTelkomTrans() error
	RekponFeeUpdateOnHalloTrans() error
}

type rekpontSchedulerUsecase struct{}

func NewRekponSchedulerUsecase() RekponSchedulerUsecase {
	return &rekpontSchedulerUsecase{}
}

func (r *rekpontSchedulerUsecase) RekponFeeUpdateOnTelkomTrans() (er error) {
	repo, _ := rekponrepo.NewRekponRepo()

	totalTrx, er := repo.CountTelkomTrans(helper.BeginCurrentDate, helper.EndCurrentDate)
	if er != nil {
		return er
	}

	data := totalTrx
	if data > 0 {
		_ = glg.Log("Scheduler INFO: ", "Update fee telkom transaction is processing..")
		totalRows, er := repo.UpdateFeeTelkomTrans(helper.BeginCurrentDate, helper.EndCurrentDate)
		if er != nil {
			return er
		}
		_ = glg.Log("Scheduler INFO: ", "Update fee telkom transaction succeeded:", totalRows, "trx")

		hours, minutes, _ := time.Now().Clock()
		currUTCTimeInString := fmt.Sprintf("%d:%02d", hours, minutes)

		_ = glg.Log("Scheduler INFO: ", "Update fee telkom transaction is done at:", currUTCTimeInString)
	} else {
		_ = glg.Log("Scheduler INFO: ", "There is no telkom transaction data to update:", data, "trx")
	}

	return
}

func (r *rekpontSchedulerUsecase) RekponFeeUpdateOnHalloTrans() (er error) {
	repo, _ := rekponrepo.NewRekponRepo()

	totalTrx, er := repo.CountHalloTrans(helper.BeginCurrentDate, helper.EndCurrentDate)
	if er != nil {
		return er
	}

	data := totalTrx
	if data > 0 {
		_ = glg.Log("Scheduler INFO: ", "Update fee hallo transaction is processing..")
		totalRows, er := repo.UpdateFeeHalloTrans(helper.BeginCurrentDate, helper.EndCurrentDate)
		if er != nil {
			return er
		}
		_ = glg.Log("Scheduler INFO: ", "Update fee hallo transaction succeeded:", totalRows, "trx")

		hours, minutes, _ := time.Now().Clock()
		currUTCTimeInString := fmt.Sprintf("%d:%02d", hours, minutes)

		_ = glg.Log("Scheduler INFO: ", "Update fee hallo transaction is done at:", currUTCTimeInString)
	} else {
		_ = glg.Log("Scheduler INFO: ", "There is no hallo transaction data to update:", data, "trx")
	}

	return
}
