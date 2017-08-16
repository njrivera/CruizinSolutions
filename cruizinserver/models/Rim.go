package models

type Rim struct {
	ItemNum     int    `json:"itemnum"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Size        string `json:"size"`
	BoltPattern string `json:"boltpattern"`
	Finish      string `json:"finish"`
	Condition   string `json:"condition"`
	Price       string `json:"price"`
	Qty         int    `json:"qty"`
}
