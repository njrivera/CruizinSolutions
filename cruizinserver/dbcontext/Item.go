package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetItems() []models.Item {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetItems)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price float32
	var qty int
	var items []models.Item
	for rows.Next() {
		rows.Scan(&itemnum, &description, &price, &qty)
		items = append(items, models.Item{
			ItemNum:     itemnum,
			Description: description,
			Price:       price,
			Qty:         qty})
	}
	db.Close()

	return items
}

func CreateItem(item models.Item) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateItem)
	util.CheckErr(err)

	statement.Exec(
		item.Description,
		item.Price,
		item.Qty)
	db.Close()

	return
}

func GetItem(key int) models.Item {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetItem, key)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price float32
	var qty int
	row.Scan(&itemnum, &description, &price, &qty)
	db.Close()
	item := models.Item{
		ItemNum:     itemnum,
		Description: description,
		Price:       price,
		Qty:         qty}

	return item
}

func DeleteItem(itemnum int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteItem)
	util.CheckErr(err)

	statement.Exec(itemnum)

	return
}
