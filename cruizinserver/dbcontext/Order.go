package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func CreateOrder(order models.Order, items []models.ItemOrder) int {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateOrder)
	util.CheckErr(err)

	row, err := statement.Exec(
		order.Date,
		order.Cid,
		order.Vid,
		order.Odometer,
		order.Comments,
		order.Subtotal,
		order.Tax,
		order.Total)
	util.CheckErr(err)
	ordernum, err := row.LastInsertId()
	util.CheckErr(err)

	statement, err = db.Prepare(queries.CreateItemOrder)
	util.CheckErr(err)
	for _, item := range items {
		statement.Exec(
			ordernum,
			item.ItemNum,
			item.Qty,
			item.Amount)
	}
	db.Close()

	return int(ordernum)
}

func GetCustVehicles(key int) []models.Vehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetCustVehicles, key)
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
