package queries

const (
	GetTires      = "SELECT * FROM tires"
	CreateTire    = "INSERT INTO tires (itemnum, brand, model, size, servicedesc, warranty, condition, price, qty) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	GetTire       = "SELECT * FROM tires WHERE itemnum = ?"
	DeleteTire    = "DELETE FROM tires WHERE itemnum = ?"
	UpdateTire    = "UPDATE tires SET brand = ?, model = ?, size = ?, servicedesc = ?, warranty = ?, condition = ?, price = ?, qty = ? WHERE itemnum = ?"
	UpdateTireQty = "UPDATE tires SET qty = (qty + ?) WHERE itemnum = ?"
)
