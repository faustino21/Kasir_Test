package entity

import "time"

type Cashier struct {
	CashierID   int        `db:"cashier_id"`
	CashierName string     `db:"name"`
	Password    string     `db:"password"`
	UpdatedAt   *time.Time `db:"updated_at"`
	CreatedAt   *time.Time `db:"created_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
