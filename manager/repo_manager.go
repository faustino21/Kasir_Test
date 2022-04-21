package manager

import "Kasir_Test/repository"

type RepoManager interface {
	CashierRepo() repository.CashierRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.NewCashierRepo(r.infra.SqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
