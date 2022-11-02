package apexrepo

import (
	"database/sql"
	"errors"
	"fmt"
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
