package models

type Order struct {
	OrderNum int    `json:"ordernum"`
	Date     string `json:"date"`
	Cid      int    `json:"cid"`
	Vid      int    `json:"vid"`
	Odometer int    `json:"odometer"`
	Comments string `json:"comments"`
	Subtotal string `json:"subtotal"`
	Tax      string `json:"tax"`
	Total    string `json:"total"`
}

type ItemOrder struct {
	Id       int    `json:"id"`
	OrderNum int    `json:"ordernum"`
	ItemNum  int    `json:"itemnum"`
	Price    string `json:"price"`
	Qty      int    `json:"qty"`
	Amount   string `json:"amount"`
}

type OrderWithItems struct {
	Order Order       `json:"order"`
	Items []ItemOrder `json:"items"`
}

type OrderWithVehicle struct {
	OrderNum int    `json:"ordernum"`
	Date     string `json:"date"`
	Cid      int    `json:"cid"`
	Vid      int    `json:"vid"`
	Year     string `json:"year"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Odometer int    `json:"odometer"`
	Comments string `json:"comments"`
	Subtotal string `json:"subtotal"`
	Tax      string `json:"tax"`
	Total    string `json:"total"`
}

type ItemOrderWithDesc struct {
	ItemNum     int    `json:"itemnum"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Qty         int    `json:"qty"`
	Amount      string `json:"amount"`
}
