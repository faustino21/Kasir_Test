package entity

type Discount struct {
	Id              int    `db:"discount_id" json:"discountId,omitempty"`
	Qty             int    `db:"qty" json:"qty,omitempty"`
	DiscType        string `db:"type" json:"type,omitempty"`
	Result          int    `db:"result" json:"result"`
	ExpiredAt       int    `db:"expired_at" json:"expiredAt"`
	ExpiredAtFormat string `json:"expiredAtFormat,omitempty"`
	StringFormat    string `json:"stringFormat,omitempty"`
}
