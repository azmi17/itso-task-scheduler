package apexrepo

import (
	"database/sql"
	"fmt"
	"itso-task-scheduler/helper"
	"testing"
	"time"
)

func GetConnectionApx() *sql.DB {
	dataSource := "root:azmic0ps@tcp(localhost:3317)/integrasi_apex_ems?parseTime=true"
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
func TestCleanUpTriggerReversal(t *testing.T) {
	// db := GetConnectionApx()
	// apexrepo := newApexRepoMysqlImpl(db)

	// err := apexrepo.CleanUpTriggerByReversalOnTabtrans()
	// if err != nil {
	// 	_ = glg.Log(err.Error())
	// }
	// fmt.Println("Delete transaction succeeded..")

	currentDate := helper.BeginCurrentDate
	fmt.Println(currentDate)

	beforeDate := helper.EndCurrentDate
	fmt.Println(beforeDate)
}
