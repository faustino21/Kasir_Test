package repository

import (
	"Kasir_Test/entity"
	"Kasir_Test/util"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/leekchan/accounting"
)

type ProductRepo interface {
	GetListProduct(limit, skip, id int, q string) (int, *[]entity.Product, error)
}

type productRepoImpl struct {
	db *sqlx.DB
}

func (p *productRepoImpl) GetListProduct(limit, skip, id int, q string) (int, *[]entity.Product, error) {
	funcName := "ProductRepo.GetListProduct"
	var productList []entity.Product
	query := "SELECT product_id, category_id, sku, name, stock, price, image, discount FROM product"
	if id != 0 && q != "" {
		err := p.db.Select(&productList, query+" WHERE category_id = ? AND name = ? LIMIT ? OFFSET ?", id,
			q, limit, skip)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
	} else if id != 0 {
		err := p.db.Select(&productList, query+" WHERE category_id = ?  LIMIT ? OFFSET ?", id, limit, skip)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
	} else if q != "" {
		err := p.db.Select(&productList, fmt.Sprintf(query+" WHERE name LIKE \"%%%s%%\" LIMIT ? OFFSET ?", q, limit, skip))
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
	} else {
		err := p.db.Select(&productList, query+" LIMIT ? OFFSET ?", limit, skip)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
	}

	for _, v := range productList {
		category, err := p.getCategory(v.CategoryId)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
		discount, err := p.getDiscount(v)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			return 0, nil, fmt.Errorf(err.Error())
		}
		v.Category = category
		v.Discount = discount
	}

	return len(productList), &productList, nil
}

func (p *productRepoImpl) getCategory(categoryId int) (*entity.Category, error) {
	funcName := "ProductRepoImpl.getCategory"
	var category entity.Category
	err := p.db.Get(&category, "SELECT * FROM category WHERE category_id = ?", categoryId)
	if err != nil {
		util.Log.Error().Msgf(funcName+" : %v", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &category, nil
}

func (p *productRepoImpl) getDiscount(product entity.Product) (*entity.Discount, error) {
	funcName := "ProductRepo.getDiscount"
	var discount entity.Discount
	var count int
	ac := accounting.Accounting{Symbol: "Rp", Precision: 0, Thousand: "."}
	err := p.db.Get(&discount, "SELECT discount_id, qty, type, result, expired_at FROM discount WHERE discount_id = ?", product.DiscountId)
	if err != sql.ErrNoRows {
		util.Log.Error().Msgf(funcName+".qureying : %v", err)
		return nil, fmt.Errorf(err.Error())
	}
	//date, err := time.Parse(time.RFC3339, discount.ExpiredAt.String())
	//if err != nil {
	//	util.Log.Error().Msgf(funcName+".parsingTime : %v", err)
	//	return nil, fmt.Errorf(err.Error())
	//}
	waktu := util.TimeUnix(discount.ExpiredAt)
	if discount.DiscType == "PERCENT" {
		count = product.Price - (product.Price * (discount.Result / 100))
	}
	stringFormat := fmt.Sprintf("Discount %d%% %s", count, ac.FormatMoney(count))
	discount.StringFormat = stringFormat
	discount.ExpiredAtFormat = waktu.Format("02 Jan 2006")
	return &discount, nil
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepoImpl{db}
}

//"02 Jan 2006"
