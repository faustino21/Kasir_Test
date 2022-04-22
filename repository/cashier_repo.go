package repository

import (
	"Kasir_Test/Delivery/httpResp"
	"Kasir_Test/entity"
	"Kasir_Test/util"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type CashierRepo interface {
	GetAll(limit, skip int) (int, *[]httpResp.ListCashier, error)
	Get(cashierId int) (*entity.Cashier, error)
	Insert(name, passcode string) (*entity.Cashier, error)
	Update(id int, name, passcode string) error
	Delete(id int) error
}

type CashierRepoImpl struct {
	db *sqlx.DB
}

func (c *CashierRepoImpl) GetAll(limit, skip int) (int, *[]httpResp.ListCashier, error) {
	funcName := "CashierRepo.GetAll"

	var cashier []httpResp.ListCashier
	var total int
	err := c.db.Select(&cashier, fmt.Sprintf("SELECT cashier_id, name FROM cashier WHERE deleted_at IS NULL LIMIT %d OFFSET %d", limit, skip))
	if err != nil {
		util.Log.Error().Msgf(funcName+" : %w", err)
		return 0, nil, errors.New(err.Error())
	}

	err = c.db.Get(&total, fmt.Sprintf("SELECT COUNT(cashier_id) FROM cashier"))
	if err != nil {
		util.Log.Error().Msgf(funcName+" : %w", err)
		return 0, nil, errors.New(err.Error())
	}

	return total, &cashier, nil
}

func (c *CashierRepoImpl) Get(cashierId int) (*entity.Cashier, error) {
	funcName := "CashierRepo.Get"

	var cashier entity.Cashier
	err := c.db.Get(&cashier, "SELECT cashier_id, name FROM cashier WHERE cashier_id = ?", cashierId)
	if err != nil {
		util.Log.Error().Msgf(funcName+" : %w", err)
		return nil, errors.New(err.Error())
	}
	return &cashier, nil
}

func (c *CashierRepoImpl) Insert(name, passcode string) (*entity.Cashier, error) {
	funcName := "CashierRepo.Insert"
	var cashier entity.Cashier

	tx := c.db.MustBegin()
	row := tx.MustExec(fmt.Sprintf("INSERT INTO cashier (name, password) VALUES (\"%s\", \"%s\")", name, passcode))
	rowAffected, err := row.RowsAffected()
	if rowAffected == 0 && err != nil {
		util.Log.Error().Msgf(funcName+".rowsAffected : %w", err)
		return nil, errors.New(err.Error())
	}

	id, err := row.LastInsertId()
	err = tx.Get(&cashier, "SELECT * FROM cashier WHERE cashier_id = ?", id)
	if err != nil {
		util.Log.Error().Msgf(funcName+".lasInsert : %w", err)
		return nil, errors.New(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		util.Log.Error().Msgf(funcName+".commit : %w", err)
		return nil, errors.New(err.Error())
	}
	return &cashier, nil
}

func (c *CashierRepoImpl) Update(id int, name, passcode string) error {
	funcName := "CashierRepo.Update"

	tx := c.db.MustBegin()
	if name != "" {
		tx.MustExec(fmt.Sprintf("UPDATE cashier SET name = \"%s\" WHERE cashier_id = %d", name, id))
	}
	if passcode != "" {
		tx.MustExec(fmt.Sprintf("UPDATE cashier SET password = \"%s\" WHERE cashier_id = %d", passcode, id))
	}
	err := tx.Commit()
	if err != nil {
		util.Log.Error().Msgf(funcName+".commit : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (c *CashierRepoImpl) Delete(id int) error {
	funcName := "CashierRepo.Delete"

	timeStamp := time.Now().Local().Format("2006-01-02 15:04:05")

	tx := c.db.MustBegin()
	tx.MustExec(fmt.Sprintf("UPDATE cashier SET deleted_at = \"%v\" WHERE cashier_id = %d", timeStamp, id))
	err := tx.Commit()
	if err != nil {
		util.Log.Error().Msgf(funcName+".commit : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewCashierRepo(db *sqlx.DB) CashierRepo {
	return &CashierRepoImpl{
		db: db,
	}
}
