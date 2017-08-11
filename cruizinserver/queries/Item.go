package queries

const (
	GetItems   = "SELECT * FROM items"
	CreateItem = "INSERT INTO items (description, price) VALUES (?, ?)"
	GetItem    = "SELECT * FROM items WHERE itemnum = ?"
	DeleteItem = "DELETE FROM items WHERE itemnum = ?"
)
