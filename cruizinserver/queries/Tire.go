package queries

const (
	GetTires   = "SELECT * FROM tires"
	CreateTire = "INSERT INTO tires (brand, model, size, servicedesc, notes) VALUES (?, ?, ?, ?, ?)"
	GetTire    = "SELECT * FROM tires WHERE partnum = ?"
	DeleteTire = "DELETE FROM tires WHERE itemnum = ?"
)
