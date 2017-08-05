package models

type Service struct {
	PartNum int     `json:"partnum"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}
