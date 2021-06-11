package dbrepo

import (
	"database/sql"

	"github.com/rzolm/DAC_Application/pkg/config"
)

	//*this is used for database driver flexibility 

type mysqlDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewMysqlRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &mysqlDBRepo
	App: a,
	BD: conn,
}