package entity

import (
	"database/sql"
	"time"
)

type Cashier struct {
	CashierID   int             `db:"cashier_id" json:"cashier_id,omitempty"`
	CashierName string          `db:"name" json:"cashier_name,omitempty"`
	Password    string          `db:"password" json:"password,omitempty"`
	UpdatedAt   *time.Time      `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt   *time.Time      `db:"created_at" json:"created_at,omitempty"`
	DeletedAt   *time.Time      `db:"deleted_at" json:"deleted_at,omitempty"`
	Token       *sql.NullString `db:"token" json:"token,omitempty"`
}
