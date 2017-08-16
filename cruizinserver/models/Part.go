package models

type Part struct {
	ItemNum     int     `json:"itemnum"`
	Description string  `json:"description"`
	Condition   string  `json:"condition"`
	Price       string `json:"price"`
}
