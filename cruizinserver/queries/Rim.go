package queries

const (
	GetRims      = "SELECT * FROM rims"
	CreateRim    = "INSERT INTO rims (itemnum, brand, model, size, boltpattern, finish, condition, price, qty) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	GetRim       = "SELECT * FROM rims WHERE itemnum = ?"
	DeleteRim    = "DELETE FROM rims WHERE itemnum = ?"
	UpdateRim    = "UPDATE rims SET brand = ?, model = ?, size = ?, boltpattern = ?, finish = ?, condition = ?, price = ?, qty = ? WHERE itemnum = ?"
	UpdateRimQty = "UPDATE rims SET qty = (qty + ?) WHERE itemnum = ?"
)
