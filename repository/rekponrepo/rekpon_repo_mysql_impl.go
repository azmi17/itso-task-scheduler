package rekponrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/entities/err"
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

// ===================================================================================================================================================
// ==================================================================Production Below=================================================================
// ===================================================================================================================================================

func (r *RekponRepoMysqlImpl) FindEmptyFeeTelkomHalloTrans(beginCurrentDate, endCurrentDate string) (data []entities.TransHistory, er error) {
	rows, er := r.rekponDb.Query(`SELECT
		th.trans_id,
		th.stan,
		th.rek_id,
		r.bank_code,
		th.biller_code,
		th.product_code,
		th.subscriber_id,
		th.dc,
		th.Response_code,
		th.Amount,
		th.profit_included,
		th.profit_excluded,
		th.profit_share_biller,
		th.Profit_share_aggregator,	
		th.profit_share_bank,
		th.markup_total,
		th.markup_share_aggregator,
		th.Markup_Share_Bank
	FROM trans_history AS th INNER JOIN rekening AS r ON (th.rek_id = r.rek_id)  
		WHERE 
		th.response_code='0000'
		AND th.dc='d'
		AND th.biller_code ='001001'
		AND th.profit_included = 0
		AND th.profit_share_biller = 0
		AND th.profit_share_aggregator = 0
		AND th.profit_share_bank = 0 
		AND th.tgl_trans_str >= ?
		AND th.tgl_trans_str <= ?`, beginCurrentDate, endCurrentDate)
	if er != nil {
		return data, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var TransHistory entities.TransHistory
		if er = rows.Scan(
			&TransHistory.TransId,
			&TransHistory.Stan,
			&TransHistory.Rek_Id,
			&TransHistory.Bank_Code,
			&TransHistory.Biller_Code,
			&TransHistory.Product_Code,
			&TransHistory.Subscriber_Id,
			&TransHistory.Dc,
			&TransHistory.Response_Code,
			&TransHistory.Amount,
			&TransHistory.Profit_Included,
			&TransHistory.Profit_Excluded,
			&TransHistory.Profit_Share_Biller,
			&TransHistory.Profit_Share_aggr,
			&TransHistory.Profitt_Share_Bank,
			&TransHistory.Markup_Total,
			&TransHistory.Markup_Share_Aggregator,
			&TransHistory.Markup_Share_Bank,
		); er != nil {
			return data, er
		}

		data = append(data, TransHistory)
	}

	if len(data) == 0 {
		return data, err.NoRecord
	} else {
		return
	}
}
func (r *RekponRepoMysqlImpl) GetFeeOnProductConfig(bankCode, billerCode, productCode string) (feeDetail entities.ProductConfig, er error) {
	row := r.rekponDb.QueryRow(`SELECT 
		bank_code,
		product_code,
		dc,
		deskripsi,
		profit_excluded,
		profit_included,
		profit_share_biller,
		profit_share_aggregator,
		profit_share_bank
	FROM product_config WHERE bank_code = ? AND biller_code = ? AND product_code = ?`,
		bankCode,
		billerCode,
		productCode)
	er = row.Scan(
		&feeDetail.BankCode,
		&feeDetail.Product_Code,
		&feeDetail.Dc,
		&feeDetail.Deskripsi,
		&feeDetail.Profit_Excluded,
		&feeDetail.Profit_Included,
		&feeDetail.Profit_Share_Biller,
		&feeDetail.Profit_Share_aggr,
		&feeDetail.Profitt_Share_Bank,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return feeDetail, err.NoRecord
		} else {
			return feeDetail, errors.New(fmt.Sprint("error while get product config: ", er.Error()))
		}
	}
	return
}

func (r *RekponRepoMysqlImpl) UpdateFeeTelkomHalloTrans(profitIncluded, profitShareBiller, profitShareAgg, profitShareBank int64, stan string) (er error) {
	stmt, er := r.rekponDb.Prepare(`UPDATE trans_history SET 
		profit_included = ?, 
		profit_share_biller = ?,
		profit_share_aggregator= ?, 
		profit_share_bank= ?
	WHERE 
		response_code='0000'
		AND dc='d'
		AND biller_code ='001001'
		AND profit_included = 0
		AND profit_share_biller = 0
		AND profit_share_aggregator = 0
		AND profit_share_bank = 0 
		AND stan = ?
	`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update fee transaction: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er = stmt.Exec(
		profitIncluded,
		profitShareBiller,
		profitShareAgg,
		profitShareBank,
		stan); er != nil {
		return errors.New(fmt.Sprint("error while update fee transaction: ", er.Error()))
	}

	return
}
