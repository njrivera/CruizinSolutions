package models

type Tire struct {
	PartNum     int     `json:"partnum"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Size        string  `json:"size"`
	SpeedRating string  `json:"speedrating"`
	LoadRange   string  `json:"loadrange"`
	Price       float32 `json:"price"`
	Qty         int     `json:"qty"`
}
