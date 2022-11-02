package rekponrepo

type RekponRepo interface {
	CountTelkomTrans(string, string) (int64, error)
	CountHalloTrans(string, string) (int64, error)
	UpdateFeeTelkomTrans(string, string) (int64, error)
	UpdateFeeHalloTrans(string, string) (int64, error)
}
