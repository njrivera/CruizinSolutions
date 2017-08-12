package models

type Part struct {
	ItemNum     int     `json:"itemnum"`
	Description string  `json:"description"`
	Condition   string  `json:"condition"`
	Price       float32 `json:"price"`
}
