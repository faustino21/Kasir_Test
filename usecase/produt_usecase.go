package usecase

import (
	"Kasir_Test/entity"
	"Kasir_Test/repository"
)

type ProductUseCase interface {
	GetAllProduct(limit, skip, id int, q string) (int, *[]entity.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepo
}

func (p *productUseCase) GetAllProduct(limit, skip, id int, q string) (int, *[]entity.Product, error) {
	return p.repo.GetListProduct(limit, skip, id, q)
}

func NewProductUseCase(repo repository.ProductRepo) ProductUseCase {
	return &productUseCase{repo}
}
