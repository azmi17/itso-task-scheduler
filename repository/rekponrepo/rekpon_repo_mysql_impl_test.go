package rekponrepo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/kpango/glg"
)

func GetConnectionRekpon() *sql.DB {
	dataSource := "root:azmic0ps@tcp(localhost:3317)/rekpon?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
func TestCountTelkomTrans(t *testing.T) {
	db := GetConnectionRekpon()
	rekponRepo := newRekponRepoMysqlImpl(db)

	total, err := rekponRepo.CountTelkomTrans("20221102000000", "20221102595959")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Total Transaction:", total)
}

func TestCountHalloTrans(t *testing.T) {
	db := GetConnectionRekpon()
	rekponRepo := newRekponRepoMysqlImpl(db)

	total, err := rekponRepo.CountHalloTrans("20221102000000", "20221102595959")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Total Transaction:", total)
}

func TestUpdateFeeTelkomTrans(t *testing.T) {
	db := GetConnectionRekpon()
	rekponRepo := newRekponRepoMysqlImpl(db)

	total, err := rekponRepo.UpdateFeeTelkomTrans("20221102000000", "20221102595958")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Total Update Transaction:", total)
}

func TestUpdateFeeTelkomHalloTrans(t *testing.T) {
	var er error

	db := GetConnectionRekpon()
	rekponRepo := newRekponRepoMysqlImpl(db)

	trans, er := rekponRepo.FindEmptyFeeTelkomHalloTrans("20221102000000", "20221102595958")
	if er != nil {
		t.Error(er.Error())
	}

	for _, val := range trans {
		feeData, er := rekponRepo.GetFeeOnProductConfig(val.Bank_Code, val.Biller_Code, val.Product_Code)
		if er != nil {
			_ = glg.Log("no records found")
		}

		feeData, er = rekponRepo.GetFeeOnProductConfig("default", val.Biller_Code, val.Product_Code)
		if er != nil {
			_ = glg.Log("no records found")
		}

		er = rekponRepo.UpdateFeeTelkomHalloTrans(int64(feeData.Profit_Included), int64(feeData.Profit_Share_Biller), int64(feeData.Profit_Share_aggr), int64(feeData.Profitt_Share_Bank), val.Stan)
		if er != nil {
			t.Error(er.Error())
		}
		fmt.Println("update fee successfully on stan:", val.Stan)
	}
}
