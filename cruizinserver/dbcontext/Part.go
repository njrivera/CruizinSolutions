package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetParts() []models.Part {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetParts)
	util.CheckErr(err)

	var itemnum int
	var description string
	var condition string
	var price string
	var parts []models.Part
	for rows.Next() {
		rows.Scan(&itemnum, &description, &condition, &price)
		parts = append(parts, models.Part{
			ItemNum:     itemnum,
			Description: description,
			Condition:   condition,
			Price:       price})
	}
	db.Close()

	return parts
}

func CreatePart(part models.Part) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreatePart)
	util.CheckErr(err)

	statement.Exec(
		part.ItemNum,
		part.Description,
		part.Condition,
		part.Price)
	db.Close()

	return
}

func GetPart(key int) models.Part {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetPart, key)
	util.CheckErr(err)

	var itemnum int
	var description string
	var condition string
	var price string
	row.Scan(&itemnum, &description, &condition, &price)
	db.Close()
	part := models.Part{
		ItemNum:     itemnum,
		Description: description,
		Condition:   condition,
		Price:       price}

	return part
}

func DeletePart(itemnum int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeletePart)
	util.CheckErr(err)

	statement.Exec(itemnum)

	return
}

func UpdatePart(part models.Part) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdatePart)
	util.CheckErr(err)

	statement.Exec(
		part.Description,
		part.Condition,
		part.Price,
		part.ItemNum)
	db.Close()

	return
}
