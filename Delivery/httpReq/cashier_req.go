package httpReq

type CashierReq struct {
	Name     string `json:"name" binding:"required"`
	Passcode string `json:"passcode" binding:"required"`
}
