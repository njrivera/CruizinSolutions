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
	var items []models.Item
	for rows.Next() {
		rows.Scan(&itemnum, &description, &price)
		items = append(items, models.Item{
			ItemNum:     itemnum,
			Description: description,
			Price:       price})
	}
	db.Close()

	return items
}

func CreateItem(item models.Item) int {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateItem)
	util.CheckErr(err)

	row, err := statement.Exec(
		item.Description,
		item.Price)
	util.CheckErr(err)
	id, err := row.LastInsertId()
	util.CheckErr(err)
	db.Close()

	return int(id)
}

func GetItem(key int) models.Item {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetItem, key)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price float32
	row.Scan(&itemnum, &description, &price)
	db.Close()
	item := models.Item{
		ItemNum:     itemnum,
		Description: description,
		Price:       price}

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
