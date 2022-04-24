package usecase

import (
	"Kasir_Test/entity"
	"Kasir_Test/repository"
)

type LoginUseCase interface {
	GetPasscode(id int) (*entity.Cashier, error)
	Verify(id int, passcode string) (*entity.Cashier, error)
	InsertingToken(id int, token string) error
	Authorize(token string) error
}

type loginUseCase struct {
	repo repository.LoginRepo
}

func (l *loginUseCase) GetPasscode(id int) (*entity.Cashier, error) {
	return l.repo.GetPasscode(id)
}

func (l *loginUseCase) Verify(id int, passcode string) (*entity.Cashier, error) {
	return l.repo.CheckCashier(id, passcode)
}

func (l *loginUseCase) InsertingToken(id int, token string) error {
	return l.repo.UpdateToken(id, token)
}

func (l *loginUseCase) Authorize(token string) error {
	return l.repo.CheckToken(token)
}

func NewLoginUseCase(repo repository.LoginRepo) LoginUseCase {
	return &loginUseCase{
		repo: repo,
	}
}
