package models

type Customer struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Middle    string `json:"middle"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
}
