package usecase

import (
	"Kasir_Test/Delivery/httpResp"
	"Kasir_Test/entity"
	"Kasir_Test/repository"
)

type CashierUseCase interface {
	GetCashier(skip, limit int) (int, *[]httpResp.ListCashier, error)
	GetCashierDetail(cashierId int) (*entity.Cashier, error)
	RegisterCashier(name, passcode string) (*entity.Cashier, error)
	UpdateCashier(id int, name, passcode string) error
	DeleteCashier(id int) error
}

type cashierUseCase struct {
	repo repository.CashierRepo
}

func (c *cashierUseCase) GetCashier(skip, limit int) (int, *[]httpResp.ListCashier, error) {
	return c.repo.GetAll(limit, skip)
}

func (c *cashierUseCase) GetCashierDetail(cashierId int) (*entity.Cashier, error) {
	return c.repo.Get(cashierId)
}

func (c *cashierUseCase) RegisterCashier(name, passcode string) (*entity.Cashier, error) {
	return c.repo.Insert(name, passcode)
}

func (c *cashierUseCase) UpdateCashier(id int, name, passcode string) error {
	return c.repo.Update(id, name, passcode)
}

func (c *cashierUseCase) DeleteCashier(id int) error {
	return c.repo.Delete(id)
}

func NewCashierUseCas(repo repository.CashierRepo) CashierUseCase {
	return &cashierUseCase{
		repo: repo,
	}
}
