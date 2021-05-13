package queries

const (
	GetParts   = "SELECT * FROM parts"
	CreatePart = "INSERT INTO parts (itemnum, description, condition, price) VALUES (?, ?, ?, ?)"
	GetPart    = "SELECT * FROM parts WHERE itemnum = ?"
	DeletePart = "DELETE FROM parts WHERE itemnum = ?"
	UpdatePart = "UPDATE parts SET description = ?, condition = ?, price = ? WHERE itemnum = ?"
)
