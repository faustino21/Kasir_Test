package entity

import "database/sql"

type Product struct {
	Id         int           `db:"product_id" json:"productId"`
	Sku        string        `db:"sku" json:"sku"`
	Name       string        `db:"name" json:"name"`
	Stock      int           `db:"stock" json:"stock"`
	Price      int           `db:"price" json:"price"`
	Image      string        `db:"image" json:"image"`
	CategoryId int           `db:"category_id"`
	DiscountId sql.NullInt32 `db:"discount"`
	Category   *Category     `json:"category"`
	Discount   *Discount     `json:"discount"`
}
