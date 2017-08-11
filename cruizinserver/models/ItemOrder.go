package models

type ItemOrder struct {
	Id       int `json:"id"`
	OrderNum int `json:"ordernum"`
	ItemNum  int `json:"itemnum"`
	Qty      int `json:"qty"`
	Subtotal int `json:"subtotal"`
}
