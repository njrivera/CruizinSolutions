package queries

const (
	GetTires   = "SELECT * FROM tires"
	GetTire    = "SELECT * FROM tires WHERE partnum = ?"
	CreateTire = "INSERT INTO tires (brand, model, size, speedrating, loadrange, price, qty) VALUES (?, ?, ?, ?, ?, ?, ?)"
)
