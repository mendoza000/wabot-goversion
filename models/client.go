package models

type Client struct {
	Name               string  `json:"name"`
	Day                int     `json:"day"`
	Screen             int     `json:"screen"`
	Paid               bool    `json:"paid"`
	Amount             float64 `json:"amount"`
	AccountID          int     `json:"account_id"`
	AccountService     int     `json:"account_service"`
	Phone              string  `json:"phone"`
	IsResellerCustomer bool    `json:"is_reseller_customer"`
}
