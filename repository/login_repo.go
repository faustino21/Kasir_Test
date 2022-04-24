package repository

import (
	"Kasir_Test/entity"
	"Kasir_Test/util"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LoginRepo interface {
	GetPasscode(id int) (*entity.Cashier, error)
	CheckCashier(id int, passcode string) (*entity.Cashier, error)
	UpdateToken(id int, token string) error
	CheckToken(token string) error
}

type LoginRepoImpl struct {
	db *sqlx.DB
}

func (l *LoginRepoImpl) GetPasscode(id int) (*entity.Cashier, error) {
	funcName := "LoginRepo.Passcode"
	var passcode entity.Cashier
	err := l.db.Get(&passcode, "SELECT password FROM cashier WHERE cashier_id = ? AND deleted_at IS NULL", id)
	if passcode.Password == "" {
		return nil, fmt.Errorf("account not found")
	}
	if err != nil {
		util.Log.Error().Msgf(funcName+".GetPassCode : %v", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &passcode, nil
}

func (l *LoginRepoImpl) CheckCashier(id int, passcode string) (*entity.Cashier, error) {
	funcName := "LoginRepo.CheckCashier"
	var cashier entity.Cashier
	err := l.db.Get(&cashier, "SELECT name, cashier_id, password, created_at FROM cashier WHERE cashier_id = ?", id)
	if err != nil {
		util.Log.Error().Msgf(funcName+".QueryError : %v", err)
		return nil, fmt.Errorf("cashier not found")
	}
	if passcode != cashier.Password {
		return nil, fmt.Errorf("passcode does not match")
	}

	return &cashier, nil
}

func (l *LoginRepoImpl) UpdateToken(id int, token string) error {
	funcName := "LoginRepo.InsertToken"

	tx := l.db.MustBegin()
	if token == "" {
		row := tx.MustExec(fmt.Sprintf("UPDATE cashier SET token = NULL WHERE cashier_id = %d", id))
		if rowAff, _ := row.RowsAffected(); rowAff == 0 {
			util.Log.Error().Msgf(funcName + "no rows affected")
			return fmt.Errorf("no rows affected")
		}
		err := tx.Commit()
		if err != nil {
			util.Log.Error().Msgf(funcName+".commit : %v", err)
			return fmt.Errorf(err.Error())
		}
		return nil
	}
	row := tx.MustExec(fmt.Sprintf("UPDATE cashier SET token = \"%s\" WHERE cashier_id = %d", token, id))
	if rowAff, _ := row.RowsAffected(); rowAff == 0 {
		util.Log.Error().Msgf(funcName + "no rows affected")
		return fmt.Errorf("no rows affected")
	}
	err := tx.Commit()
	if err != nil {
		util.Log.Error().Msgf(funcName+".commit : %v", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (l *LoginRepoImpl) CheckToken(token string) error {
	funcName := "LoginRepo.CheckToken"

	var dbToken *entity.Cashier
	err := l.db.Get(&dbToken, "SELECT token FROM cashier WHERE token = ?", token)
	if err != nil {
		util.Log.Error().Msgf(funcName+".queryToken : %v", err)
		return fmt.Errorf(err.Error())
	}
	if dbToken.Token == nil {
		return fmt.Errorf("Unauthorized")
	}
	return nil
}

func NewLoginRepo(db *sqlx.DB) LoginRepo {
	return &LoginRepoImpl{
		db: db,
	}
}
