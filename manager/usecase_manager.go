package manager

import "Kasir_Test/usecase"

type UseCaseManager interface {
	CashierUseCase() usecase.CashierUseCase
	LoginUseCase() usecase.LoginUseCase
	ProductUseCase() usecase.ProductUseCase
}

type useCaseManager struct {
	repoMag RepoManager
}

func (u *useCaseManager) CashierUseCase() usecase.CashierUseCase {
	return usecase.NewCashierUseCas(u.repoMag.CashierRepo())
}

func (u *useCaseManager) LoginUseCase() usecase.LoginUseCase {
	return usecase.NewLoginUseCase(u.repoMag.LoginRepo())
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoMag.ProductRepo())
}

func NewUseCaseManager(repoMag RepoManager) UseCaseManager {
	return &useCaseManager{
		repoMag: repoMag,
	}
}
