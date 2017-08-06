package queries

const (
	GetVehicles   = "SELECT * FROM vehicles"
	CreateVehicle = "INSERT INTO vehicles (year, make, model) VALUES (?, ?, ?)"
	GetVehicle    = "SELECT * FROM vehicles WHERE id = ?"
	DeleteVehicle = "DELETE FROM vehicles WHERE vid = ?"
)
