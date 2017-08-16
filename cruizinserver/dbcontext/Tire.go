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
	var condition string
	var price string
	var qty int
	var tires []models.Tire
	for rows.Next() {
		rows.Scan(&itemnum, &brand, &model, &size, &servicedesc, &condition, &price, &qty)
		tires = append(tires, models.Tire{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			ServiceDesc: servicedesc,
			Condition:   condition,
			Price:       price,
			Qty:         qty})
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
		tire.ItemNum,
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.ServiceDesc,
		tire.Condition,
		tire.Price,
		tire.Qty)
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
	var condition string
	var price string
	var qty int
	row.Scan(&itemnum, &brand, &model, &size, &servicedesc, &condition, &price, &qty)
	db.Close()
	tire := models.Tire{
		ItemNum:     itemnum,
		Brand:       brand,
		Model:       model,
		Size:        size,
		ServiceDesc: servicedesc,
		Condition:   condition,
		Price:       price,
		Qty:         qty}

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

func UpdateTire(tire models.Tire) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateTire)
	util.CheckErr(err)

	statement.Exec(
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.ServiceDesc,
		tire.Condition,
		tire.Price,
		tire.Qty,
		tire.ItemNum)
	db.Close()

	return
}

func UpdateTireQty(itemnum int, qty int) models.Tire {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateTireQty)
	util.CheckErr(err)

	statement.Exec(qty, itemnum)
	tire := GetTire(itemnum)
	db.Close()

	return tire
}
