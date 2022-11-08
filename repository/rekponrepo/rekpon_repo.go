package rekponrepo

import "itso-task-scheduler/entities"

type RekponRepo interface {
	// Not Used:
	CountTelkomTrans(string, string) (int64, error)
	CountHalloTrans(string, string) (int64, error)
	UpdateFeeTelkomTrans(string, string) (int64, error)
	UpdateFeeHalloTrans(string, string) (int64, error)
	// ===================================================================================================================================================
	// ==================================================================Production Below=================================================================
	// ===================================================================================================================================================
	FindEmptyFeeTelkomHalloTrans(string, string) ([]entities.TransHistory, error)
	GetFeeOnProductConfig(string, string, string) (entities.ProductConfig, error)
	UpdateFeeTelkomHalloTrans(int64, int64, int64, int64, string) error
}
