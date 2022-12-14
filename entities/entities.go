package entities

const (
	SCHEDULER_ENABLE  = "ENABLE"
	SCHEDULER_DISABLE = "DISABLE"
)

type LKMlist struct {
	KodeLKM string
}

type TransHistory struct {
	TransId                 int64
	Stan                    string
	Rek_Id                  string
	Bank_Code               string
	Biller_Code             string
	Product_Code            string
	Subscriber_Id           string
	Dc                      string
	Response_Code           string
	Amount                  string
	Profit_Included         int
	Profit_Excluded         int
	Profit_Share_Biller     int
	Profit_Share_aggr       int
	Profitt_Share_Bank      int
	Markup_Total            int
	Markup_Share_Aggregator int
	Markup_Share_Bank       int
}

type ProductConfig struct {
	BankCode            string
	Biller_Code         string
	Product_Code        string
	Dc                  string
	Deskripsi           string
	Profit_Excluded     int
	Profit_Included     int
	Profit_Share_Biller int
	Profit_Share_aggr   int
	Profitt_Share_Bank  int
}

type RepostingData struct {
	KodeLKM     string
	TotalDebet  float64
	TotalKredit float64
}

type CalculateSaldoResult struct {
	KodeLKM    string
	SaldoAkhir float64
}

type SchedulerResponse struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
