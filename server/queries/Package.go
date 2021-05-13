package queries

const (
	GetPackages   = "SELECT * FROM packages"
	CreatePackage = "INSERT INTO packages (itemnum, description, price) VALUES (?, ?, ?)"
	GetPackage    = "SELECT * FROM packages WHERE itemnum = ?"
	DeletePackage = "DELETE FROM packages WHERE itemnum = ?"
	UpdatePackage = "UPDATE packages SET description = ?, price = ? WHERE itemnum = ?"
)
