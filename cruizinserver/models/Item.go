package models

type Item struct {
	ItemNum     int    `json:"itemnum"`
	Description string `json:"description"`
	Taxable     string `json:"taxable"`
}
