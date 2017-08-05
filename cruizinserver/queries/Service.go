package queries

const (
	GetServices   = "SELECT * FROM services"
	GetService    = "SELECT * FROM services WHERE partnum = ?"
	CreateService = "INSERT INTO services (name, price) VALUES (?, ?)"
)
