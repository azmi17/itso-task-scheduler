package apexrepo

import "itso-task-scheduler/entities"

type ApexRepo interface {
	CleanUpTriggerByReversalOnTabtrans() error
	GetRekeningLKMByStatusActive() ([]string, error)
	CalculateSaldoOnRekeningLKM(kodeLKM string) (entities.CalculateSaldoResult, error)
	RepostingSaldoOnRekeningLKMByScheduler(listOfKodeLKM ...string) error
	doRepostingSaldoProcs(data string) error
}
