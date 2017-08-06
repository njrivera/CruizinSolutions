package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetVehicles() []models.Vehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetVehicles)
	util.CheckErr(err)

	var vid int
	var year int
	var make string
	var model string
	var vehicles []models.Vehicle
	for rows.Next() {
		rows.Scan(&vid, &year, &make, &model)
		vehicles = append(vehicles, models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  make,
			Model: model})
	}
	db.Close()

	return vehicles
}

func CreateVehicle(vehicle models.Vehicle) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateVehicle)
	util.CheckErr(err)

	statement.Exec(
		vehicle.Year,
		vehicle.Make,
		vehicle.Model)
	db.Close()

	return
}

func GetVehicle(key int) models.Vehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetVehicle, key)
	util.CheckErr(err)

	var vid int
	var year int
	var make string
	var model string
	row.Scan(&vid, &year, &make, &model)
	db.Close()
	vehicle := models.Vehicle{
		Vid:   vid,
		Year:  year,
		Make:  make,
		Model: model}

	return vehicle
}

func DeleteVehicle(vid int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteVehicle)
	util.CheckErr(err)

	statement.Exec(vid)

	return
}
