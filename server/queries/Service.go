package queries

const (
	GetServices   = "SELECT * FROM services"
	CreateService = "INSERT INTO services (itemnum, description, price) VALUES (?, ?, ?)"
	GetService    = "SELECT * FROM services WHERE itemnum = ?"
	DeleteService = "DELETE FROM services WHERE itemnum = ?"
	UpdateService = "UPDATE services SET description = ?, price = ? WHERE itemnum = ?"
)
