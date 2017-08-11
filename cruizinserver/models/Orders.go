package models

type Order struct {
	OrderNum int     `json:"ordernum"`
	Date     string  `json:"date"`
	Cid      int     `json:"cid"`
	Vid      int     `json:"vid"`
	Odometer int     `json:"odometer"`
	Comments string  `json:"comments"`
	Subtotal float32 `json:"subtotal"`
	Tax      float32 `json:"tax"`
	Total    float32 `json:"total"`
}
