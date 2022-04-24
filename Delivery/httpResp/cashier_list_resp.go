package httpResp

type ListCashier struct {
	CashierID int    `json:"cashierId" db:"cashier_id"`
	Name      string `json:"name" db:"name"`
}
