package usecase

import (
	"fmt"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/entities/err"
	"itso-task-scheduler/helper"
	"itso-task-scheduler/repository/rekponrepo"
	"time"

	"github.com/kpango/glg"
)

type RekponSchedulerUsecase interface {
	RekponFeeUpdateOnTelkomTrans() error
	RekponFeeUpdateOnHalloTrans() error
	RekponUpdateFeeOnTelkomHalloTrans() error
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

// ===================================================================================================================================================
// ==================================================================Production Below=================================================================
// ===================================================================================================================================================

func (r *rekpontSchedulerUsecase) RekponUpdateFeeOnTelkomHalloTrans() (er error) {
	repo, _ := rekponrepo.NewRekponRepo()

	trxList, er := repo.FindEmptyFeeTelkomHalloTrans(helper.BeginCurrentDate, helper.EndCurrentDate)
	if er != nil {
		_ = glg.Log("Scheduler INFO:", "There is no fee transaction data to update..")
	}

	data := trxList
	if len(data) > 0 {
		_ = glg.Log("Update fee telkom & hallo transaction is processing..")
		for _, trans := range data {
			feeData, er := repo.GetFeeOnProductConfig(trans.Bank_Code, trans.Biller_Code, trans.Product_Code)
			if er != nil {
				if er == err.NoRecord {
					feeData, er = repo.GetFeeOnProductConfig("default", trans.Biller_Code, trans.Product_Code)
					if er != nil {
						entities.PrintError(er.Error())
						_ = glg.Log("error while get product config by default:", er.Error())
					}
				} else {
					entities.PrintError(er.Error())
					_ = glg.Log(er.Error())
					continue
				}
			}
			er = repo.UpdateFeeTelkomHalloTrans(
				int64(feeData.Profit_Included),
				int64(feeData.Profit_Share_Biller),
				int64(feeData.Profit_Share_aggr),
				int64(feeData.Profitt_Share_Bank),
				trans.Stan)
			if er != nil {
				entities.PrintError(er.Error())
				_ = glg.Log("error while execute fee transaction => ", er.Error())
			}
			_ = glg.Log("Update fee successfully on stan:", trans.Stan, "=> product code: (", trans.Product_Code, ")")

		}
		hours, minutes, _ := time.Now().Clock()
		currUTCTimeInString := fmt.Sprintf("%d:%02d", hours, minutes)

		_ = glg.Log("Scheduler INFO:", "Update fee telkom & hallo transaction is done at:", currUTCTimeInString)
	}

	return
}
