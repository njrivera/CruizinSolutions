package models

type Tire struct {
	ItemNum     int    `json:"itemnum"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Size        string `json:"size"`
	ServiceDesc string `json:"servicedesc"`
	Notes       string `json:"notes"`
}
