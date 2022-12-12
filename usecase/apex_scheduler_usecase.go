package usecase

import (
	"fmt"
	"itso-task-scheduler/repository/apexrepo"
	"time"

	"github.com/kpango/glg"
)

type ApexSchedulerUsecase interface {
	CleanUpTriggerReversalOnTabtrans() error
	RepostingSaldoApexByScheduler() (er error)
}

type apextSchedulerUsecase struct{}

func NewApexSchedulerUsecase() ApexSchedulerUsecase {
	return &apextSchedulerUsecase{}
}

func (a *apextSchedulerUsecase) CleanUpTriggerReversalOnTabtrans() (er error) {
	repo, _ := apexrepo.NewApexRepo()

	_ = glg.Log("Scheduler INFO: ", "Clean up trigger-reversal is processing..")

	er = repo.CleanUpTriggerByReversalOnTabtrans()
	if er != nil {
		return er
	}

	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d:%02d", hours, minutes)
	_ = glg.Log("Scheduler INFO: ", "Clean up trigger-reversal is done at:", currUTCTimeInString)

	return
}

func (a *apextSchedulerUsecase) RepostingSaldoApexByScheduler() (er error) {
	repo, _ := apexrepo.NewApexRepo()

	list, er := repo.GetRekeningLKMByStatusActive()
	if er != nil {
		return er
	}

	_ = glg.Log("Reposting saldo apex is processing..")
	er = repo.RepostingSaldoOnRekeningLKMByScheduler(list...)
	if er != nil {
		return er
	}

	return nil
}
