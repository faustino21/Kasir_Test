package manager

import "Kasir_Test/repository"

type RepoManager interface {
	CashierRepo() repository.CashierRepo
	LoginRepo() repository.LoginRepo
	ProductRepo() repository.ProductRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.NewCashierRepo(r.infra.SqlDb())
}

func (r *repoManager) LoginRepo() repository.LoginRepo {
	return repository.NewLoginRepo(r.infra.SqlDb())
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(r.infra.SqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
