package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetCustVehicles(key int) []models.Vehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetCustVehicles, key)
	util.CheckErr(err)

	var vid int
	var year int
	var make string
	var model string
	var trim string
	var vehicles []models.Vehicle
	for rows.Next() {
		rows.Scan(&vid, &year, &make, &model, &trim)
		vehicles = append(vehicles, models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  make,
			Model: model,
			Trim:  trim})
	}
	db.Close()

	return vehicles
}
