package models

type Customer struct {
	Cid     int    `json:"cid"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Phone   string `json:"phone"`
}
