package manager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type InfraManager interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}

func NewInfraManager(dataSourceName string) InfraManager {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: db,
	}
}
