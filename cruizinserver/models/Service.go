package models

type Service struct {
	ItemNum     int     `json:"itemnum"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
