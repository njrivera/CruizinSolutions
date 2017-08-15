package queries

const (
	GetItems   = "SELECT * FROM items"
	CreateItem = "INSERT INTO items (description) VALUES (?)"
	GetItem    = "SELECT * FROM items WHERE itemnum = ?"
	DeleteItem = "DELETE FROM items WHERE itemnum = ?"
	UpdateItem = "UPDATE items SET description = ? WHERE itemnum = ?"
)
