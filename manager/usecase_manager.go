package manager

import "Kasir_Test/usecase"

type UseCaseManager interface {
	CashierUseCase() usecase.CashierUseCase
}

type useCaseManager struct {
	repoMag RepoManager
}

func (u *useCaseManager) CashierUseCase() usecase.CashierUseCase {
	return usecase.NewCashierUseCas(u.repoMag.CashierRepo())
}

func NewUseCaseManager(repoMag RepoManager) UseCaseManager {
	return &useCaseManager{
		repoMag: repoMag,
	}
}
