package httpReq

type CashierReq struct {
	Name     string `json:"name"`
	Passcode string `json:"passcode" binding:"required"`
}
