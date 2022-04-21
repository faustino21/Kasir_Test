package repository

import "database/sql"

type CashierRepo interface {
}

type CashierRepoImpl struct {
	db *sql.DB
}

func NewCashierRepo(db *sql.DB) CashierRepo {
	return &CashierRepoImpl{
		db: db,
	}
}
