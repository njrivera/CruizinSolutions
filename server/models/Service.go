package models

type Service struct {
	ItemNum     int     `json:"itemnum"`
	Description string  `json:"description"`
	Price       string `json:"price"`
}
