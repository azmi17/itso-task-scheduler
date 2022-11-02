package rekponrepo

import (
	"database/sql"
	"errors"
	"fmt"
)

func newRekponRepoMysqlImpl(rekponConn *sql.DB) RekponRepo {
	return &RekponRepoMysqlImpl{
		rekponDb: rekponConn,
	}
}

type RekponRepoMysqlImpl struct {
	rekponDb *sql.DB
}

func (r *RekponRepoMysqlImpl) CountTelkomTrans(startTime, endTime string) (total int64, er error) {
	rows, err := r.rekponDb.Query(`SELECT COUNT(*) AS count_telkom_trans 
	FROM trans_history
	 WHERE response_code='0000'
	 AND dc='d'
	 AND biller_code ='001001'
	 AND product_code ='001001'
	 AND profit_included = 0
	 AND profit_share_biller = 0
	 AND profit_share_aggregator = 0
	 AND profit_share_bank = 0
	 AND tgl_trans_str >= ?
	 AND tgl_trans_str <= ?
	`, startTime, endTime)
	if err != nil {
		return 0, err
	} else {
		for rows.Next() {
			rows.Scan(&total)
		}
		return total, nil
	}
}

func (r *RekponRepoMysqlImpl) CountHalloTrans(startTime, endTime string) (total int64, er error) {
	rows, err := r.rekponDb.Query(`SELECT COUNT(*) AS count_hallo_trans 
	FROM trans_history
	 WHERE response_code='0000'
	 AND dc='d'
	 AND biller_code ='001001'
	 AND product_code ='010001'
	 AND profit_included = 0
	 AND profit_share_aggregator = 0
	 AND profit_share_bank = 0
	 AND tgl_trans_str >= ?
	 AND tgl_trans_str <= ?
	`, startTime, endTime)
	if err != nil {
		return 0, err
	} else {
		for rows.Next() {
			rows.Scan(&total)
		}
		return total, nil
	}
}

func (r *RekponRepoMysqlImpl) UpdateFeeTelkomTrans(startTime, endTime string) (totalAffectedRows int64, er error) {
	stmt, er := r.rekponDb.Prepare(`UPDATE trans_history SET 
		profit_included=3000, 
		profit_share_biller=1200, 
		profit_share_aggregator=300, 
		profit_share_bank=1500 
	WHERE response_code='0000'
	AND dc='d'
	AND biller_code ='001001'
	AND product_code ='001001'
	AND profit_included = 0
	AND profit_share_biller = 0
	AND profit_share_aggregator = 0
	AND profit_share_bank = 0
	AND tgl_trans_str >= ?
	AND tgl_trans_str <= ?
	`)
	if er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while prepare update fee telkom transaction: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	var res sql.Result
	if res, er = stmt.Exec(
		startTime,
		endTime); er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while update fee telkom transaction: ", er.Error()))
	}
	totalAffectedRows, er = res.RowsAffected()
	if er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while get total affected rows telkom transaction: ", er.Error()))
	}

	return totalAffectedRows, nil
}

func (r *RekponRepoMysqlImpl) UpdateFeeHalloTrans(startTime, endTime string) (totalAffectedRows int64, er error) {
	stmt, er := r.rekponDb.Prepare(`UPDATE trans_history SET 
		profit_included=2000, 
		profit_share_aggregator=500, 
		profit_share_bank=1500 
	WHERE response_code='0000'
	AND dc='d'
	AND biller_code ='001001'
	AND product_code ='010001'
	AND profit_included = 0
	AND profit_share_aggregator = 0
	AND profit_share_bank = 0
	AND tgl_trans_str >= ?
	AND tgl_trans_str <= ?
	`)
	if er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while prepare update fee hallo transaction: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	var res sql.Result
	if res, er = stmt.Exec(
		startTime,
		endTime); er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while update fee hallo transaction: ", er.Error()))
	}
	totalAffectedRows, er = res.RowsAffected()
	if er != nil {
		return totalAffectedRows, errors.New(fmt.Sprint("error while get total affected rows hallo transaction: ", er.Error()))
	}

	return totalAffectedRows, nil
}
