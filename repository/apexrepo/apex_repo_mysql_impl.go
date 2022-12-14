package apexrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/entities/err"
	"sync"
)

func newApexRepoMysqlImpl(apexConn *sql.DB) ApexRepo {
	return &ApexRepoMysqlImpl{
		apexDb: apexConn,
	}
}

type ApexRepoMysqlImpl struct {
	apexDb *sql.DB
}

func (a *ApexRepoMysqlImpl) CleanUpTriggerByReversalOnTabtrans() error {
	stmt, er := a.apexDb.Prepare(`DELETE FROM tabtrans WHERE pokok = 0 AND keterangan LIKE "%trigger-Reversal%"`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabtrans trigger-reversal data: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()
	if _, er := stmt.Exec(); er != nil {
		return errors.New(fmt.Sprint("error while delete tabtrans trigger-reversal datan: ", er.Error()))
	}
	return nil
}

func (a *ApexRepoMysqlImpl) GetRekeningLKMByStatusActive() (lists []string, er error) {
	// rows, er := a.apexDb.Query("SELECT no_rekening FROM tabung WHERE status = 1")
	rows, er := a.apexDb.Query("SELECT no_rekening FROM tabung")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var list entities.LKMlist
		if er = rows.Scan(&list.KodeLKM); er != nil {
			return lists, er
		}

		lists = append(lists, list.KodeLKM)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (a *ApexRepoMysqlImpl) CalculateSaldoOnRekeningLKM(kodeLKM string) (data entities.CalculateSaldoResult, er error) {
	var tabtrans entities.RepostingData
	rows, err := a.apexDb.Query(`SELECT
	  tab.no_rekening,
	  SUM(CASE WHEN trans.my_kode_trans='100' THEN trans.pokok ELSE 0 END) AS total_kredit,
	  SUM(CASE WHEN trans.my_kode_trans='200' THEN trans.pokok ELSE 0 END) AS total_debet
	FROM tabung AS tab LEFT JOIN tabtrans AS trans ON (tab.no_rekening = trans.no_rekening)
	WHERE tab.no_rekening = ? GROUP BY tab.no_rekening`, kodeLKM)
	if err != nil {
		return data, er
	}
	for rows.Next() {
		rows.Scan(
			&tabtrans.KodeLKM,
			&tabtrans.TotalKredit,
			&tabtrans.TotalDebet,
		)
	}
	data.KodeLKM = tabtrans.KodeLKM
	data.SaldoAkhir = tabtrans.TotalKredit - tabtrans.TotalDebet
	return data, nil
}

// func (a *ApexRepoMysqlImpl) CalculateSaldoOnRekeningLKM(kodeLKM string) (data entities.CalculateSaldoResult, er error) {
// 	var tabtrans entities.RepostingData
// 	row := a.apexDb.QueryRow(`SELECT
// 	  no_rekening,
// 	  SUM(CASE WHEN my_kode_trans='100' THEN pokok ELSE 0 END) AS total_kredit,
// 	  SUM(CASE WHEN my_kode_trans='200' THEN pokok ELSE 0 END) AS total_debet
// 	FROM tabtrans
// 	 WHERE
// 	no_rekening = ? GROUP BY no_rekening
// 	`, kodeLKM)
// 	er = row.Scan(
// 		&tabtrans.KodeLKM,
// 		&tabtrans.TotalKredit,
// 		&tabtrans.TotalDebet,
// 	)
// 	if er != nil {
// 		if er == sql.ErrNoRows {
// 			return data, err.NoRecord
// 		} else {
// 			return data, errors.New(fmt.Sprint("error while get reposting data: ", er.Error()))
// 		}

// 	}
// 	data.KodeLKM = tabtrans.KodeLKM
// 	data.SaldoAkhir = tabtrans.TotalKredit - tabtrans.TotalDebet

// 	return data, nil
// }

func (a *ApexRepoMysqlImpl) RepostingSaldoOnRekeningLKMByScheduler(listOfKodeLKM ...string) (er error) {
	var wg sync.WaitGroup

	entities.PrintRepoChan <- entities.PrintRepo{Status: entities.PRINT_INIT_REPO_CHAN, Size: len(listOfKodeLKM)}

	for _, each := range listOfKodeLKM {

		wg.Add(1)

		go func(each string, w *sync.WaitGroup) {
			defer w.Done()

			var status = entities.PRINT_SUCCESS_STATUS_REPO_CHAN
			var msg = entities.PRINT_SUCCESS_MSG_REPO_CHAN

			er := a.doRepostingSaldoProcs(each)
			if er != nil {
				status = entities.PRINT_FAILED_STATUS_REPO_CHAN
				msg = er.Error()
			}
			var printRepo = entities.PrintRepo{
				KodeLKM: each,
				Status:  status,
				Message: msg,
			}
			entities.PrintRepoChan <- printRepo
		}(each, &wg)
	}
	wg.Wait()
	entities.PrintRepoChan <- entities.PrintRepo{Status: entities.PRINT_FINISH_REPO_CHAN}
	return
}

func (a *ApexRepoMysqlImpl) doRepostingSaldoProcs(data string) (er error) {
	lkm, er := a.CalculateSaldoOnRekeningLKM(data)
	if er != nil {
		return errors.New(fmt.Sprint("error while calculating saldo: ", er.Error()))
	}
	stmt, er := a.apexDb.Prepare(`UPDATE tabung SET saldo_akhir = ? WHERE no_rekening = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare reposting saldo: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()
	if _, er = stmt.Exec(
		lkm.SaldoAkhir,
		lkm.KodeLKM,
	); er != nil {
		return errors.New(fmt.Sprint("error while processing reposting saldo: ", er.Error()))
	}
	return
}
