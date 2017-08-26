package models

type NewTireTaxReport struct {
	Qty int
	Tax string
}

type NewTire struct {
	Date   string
	Qty    int
	Amount string
}
