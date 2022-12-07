package entities

const (
	PRINTOUT_TYPE_LOG = iota
	PRINTOUT_TYPE_ERR

	PRINT_SUCCESS_STATUS_REPO_CHAN = "00"
	PRINT_FAILED_STATUS_REPO_CHAN  = "01"
	PRINT_INIT_REPO_CHAN           = "02"
	PRINT_FINISH_REPO_CHAN         = "03"
	PRINT_SUCCESS_MSG_REPO_CHAN    = "SUCCESS"
	PRINT_FAILED_MSG_REPO_CHAN     = "FAILED"
)

var (
	PrintOutChan  = make(chan PrintOut)
	PrintRepoChan = make(chan PrintRepo)
)

type PrintRepo struct {
	KodeLKM string
	Status  string
	Message string
	Size    int
}

type PrintOut struct {
	Type    int
	Message []interface{}
}

func PrintError(message ...interface{}) {
	po := PrintOut{
		Type:    PRINTOUT_TYPE_ERR,
		Message: message,
	}

	PrintOutChan <- po
}

func PrintLog(message ...interface{}) {
	po := PrintOut{
		Type:    PRINTOUT_TYPE_LOG,
		Message: message,
	}
	PrintOutChan <- po
}
