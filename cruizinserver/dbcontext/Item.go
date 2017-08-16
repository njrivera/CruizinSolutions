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
	var items []models.Item
	for rows.Next() {
		rows.Scan(&itemnum, &description)
		items = append(items, models.Item{
			ItemNum:     itemnum,
			Description: description})
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
		item.Description)
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
	row.Scan(&itemnum, &description)
	db.Close()
	item := models.Item{
		ItemNum:     itemnum,
		Description: description}

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

func UpdateItem(itemnum int, description string) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateItem)
	util.CheckErr(err)

	statement.Exec(
		description,
		itemnum)
	db.Close()

	return
}
