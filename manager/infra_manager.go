package manager

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type InfraManager interface {
	SqlDb() *sql.DB
}

type infra struct {
	db *sql.DB
}

func (i *infra) SqlDb() *sql.DB {
	return i.db
}

func NewInfraManager(dataSourceName string) InfraManager {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: db,
	}
}
