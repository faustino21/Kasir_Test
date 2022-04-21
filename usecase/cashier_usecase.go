package usecase

import "Kasir_Test/repository"

type CashierUseCase interface {
}

type cashierUseCase struct {
	repo repository.CashierRepo
}

func NewCashierUseCas(repo repository.CashierRepo) CashierUseCase {
	return &cashierUseCase{
		repo: repo,
	}
}
