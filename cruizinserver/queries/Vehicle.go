package queries

const (
	GetVehicles   = "SELECT * FROM vehicles"
	CreateVehicle = "INSERT INTO vehicles (year, make, model) VALUES (?, ?, ?)"
	GetVehicle    = "SELECT * FROM vehicles WHERE vid = ?"
	DeleteVehicle = "DELETE FROM vehicles WHERE vid = ?"
	UpdateVehicle = "UPDATE vehicles SET year = ?, make = ?, model = ? WHERE vid = ?"
)
