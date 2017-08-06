package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetTires() []models.Tire {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetTires)
	util.CheckErr(err)

	var itemnum int
	var brand string
	var model string
	var size string
	var servicedesc string
	var notes string
	var tires []models.Tire
	for rows.Next() {
		rows.Scan(&itemnum, &brand, &model, &size, &servicedesc, &notes)
		tires = append(tires, models.Tire{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			ServiceDesc: servicedesc,
			Notes:       notes})
	}
	db.Close()

	return tires
}

func CreateTire(tire models.Tire) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateTire)
	util.CheckErr(err)

	statement.Exec(
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.ServiceDesc,
		tire.Notes)
	db.Close()

	return
}

func GetTire(key int) models.Tire {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetTire, key)
	util.CheckErr(err)

	var itemnum int
	var brand string
	var model string
	var size string
	var servicedesc string
	var notes string
	row.Scan(&itemnum, &brand, &model, &size, &servicedesc, &notes)
	db.Close()
	tire := models.Tire{
		ItemNum:     itemnum,
		Brand:       brand,
		Model:       model,
		Size:        size,
		ServiceDesc: servicedesc,
		Notes:       notes}

	return tire
}

func DeleteTire(itemnum int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteTire)
	util.CheckErr(err)

	statement.Exec(itemnum)

	return
}
