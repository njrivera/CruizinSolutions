package queries

const (
	GetTires   = "SELECT * FROM tires"
	CreateTire = "INSERT INTO tires (itemnum, brand, model, size, servicedesc, notes, price, qty) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	GetTire    = "SELECT * FROM tires WHERE itemnum = ?"
	DeleteTire = "DELETE FROM tires WHERE itemnum = ?"
	UpdateTire = "UPDATE tires SET brand = ?, model = ?, size = ?, servicedesc = ?, notes = ?, price = ?, qty = ? WHERE itemnum = ?"
)
