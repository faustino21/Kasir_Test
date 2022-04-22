package httpResp

type ListCashier struct {
	CashierID int    `json:"cashierId" db:"cashier_id"`
	Name      string `json:"name" db:"name"`
}

func NewListCashier(id int, name string) ListCashier {
	return ListCashier{
		CashierID: id,
		Name:      name,
	}
}
