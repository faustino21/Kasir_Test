package entity

type Category struct {
	Id   int    `db:"category_id" json:"categoryId,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
}
