package queries

const (
	GetVehicles   = "SELECT * FROM vehicles"
	GetVehicle    = "SELECT * FROM vehicles WHERE id = ?"
	CreateVehicle = "INSERT INTO vehicles (year, make, model, trim) VALUES (?, ?, ?, ?)"
)
