package models

type clientx struct {
	Screen   string `json:"screen"`
	ClientID string `json:"client_id"`
}

type Account struct {
	ID         int       `json:"id"`
	CreatedAt  string    `json:"created_at"`
	Mail       string    `json:"mail"`
	Password   string    `json:"password"`
	Paid       bool      `json:"paid"`
	ServiceID  int       `json:"service_id"`
	Clients    []int     `json:"clients"`
	Amount     float64   `json:"amount"`
	MaxClients int       `json:"max_clients"`
	PayDay     int       `json:"pay_day"`
	Clientsx   []clientx `json:"clientsx"`
}
