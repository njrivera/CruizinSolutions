package models

type Item struct {
	ItemNum     int     `json:"itemnum"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
