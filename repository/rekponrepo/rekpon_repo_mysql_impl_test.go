package rekponrepo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestUpdateFeeHalloTrans(t *testing.T) {
	db := GetConnectionRekpon()
	rekponRepo := newRekponRepoMysqlImpl(db)

	total, err := rekponRepo.UpdateFeeHalloTrans("20221102000000", "20221102595958")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Total Update Transaction:", total)
}
