package apexrepo

import (
	"database/sql"
	"errors"
	"itso-task-scheduler/repository/databasefactory"
	"itso-task-scheduler/repository/databasefactory/drivers"
)

func NewApexRepo() (ApexRepo, error) {
	apexConn := databasefactory.Apex.GetConnection()
	currentDriver := databasefactory.Apex.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newApexRepoMysqlImpl(apexConn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}
}
