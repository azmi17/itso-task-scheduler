package rekponrepo

import (
	"database/sql"
	"errors"
	"itso-task-scheduler/repository/databasefactory"
	"itso-task-scheduler/repository/databasefactory/drivers"
)

func NewRekponRepo() (RekponRepo, error) {
	rekponConn := databasefactory.Rekpon.GetConnection()
	currentDriver := databasefactory.Rekpon.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newRekponRepoMysqlImpl(rekponConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
