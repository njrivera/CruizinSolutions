package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetRims() []models.Rim {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetRims)
	util.CheckErr(err)

	var itemnum int
	var brand string
	var model string
	var size string
	var boltpattern string
	var finish string
	var condition string
	var price string
	var qty int
	var rims []models.Rim
	for rows.Next() {
		rows.Scan(&itemnum, &brand, &model, &size, &boltpattern, &finish, &condition, &price, &qty)
		rims = append(rims, models.Rim{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			BoltPattern: boltpattern,
			Finish:      finish,
			Condition:   condition,
			Price:       price,
			Qty:         qty})
	}
	db.Close()

	return rims
}

func CreateRim(rim models.Rim) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateRim)
	util.CheckErr(err)

	statement.Exec(
		rim.ItemNum,
		rim.Brand,
		rim.Model,
		rim.Size,
		rim.BoltPattern,
		rim.Finish,
		rim.Condition,
		rim.Price,
		rim.Qty)
	db.Close()

	return
}

func GetRim(key int) models.Rim {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetRim, key)
	util.CheckErr(err)

	var itemnum int
	var brand string
	var model string
	var size string
	var boltpattern string
	var finish string
	var condition string
	var price string
	var qty int
	row.Scan(&itemnum, &brand, &model, &size, &boltpattern, &finish, &condition, &price, &qty)
	db.Close()
	rim := models.Rim{
		ItemNum:     itemnum,
		Brand:       brand,
		Model:       model,
		Size:        size,
		BoltPattern: boltpattern,
		Finish:      finish,
		Condition:   condition,
		Price:       price,
		Qty:         qty}

	return rim
}

func DeleteRim(itemnum int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteRim)
	util.CheckErr(err)

	statement.Exec(itemnum)

	return
}

func UpdateRim(rim models.Rim) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateRim)
	util.CheckErr(err)

	statement.Exec(
		rim.Brand,
		rim.Model,
		rim.Size,
		rim.BoltPattern,
		rim.Finish,
		rim.Condition,
		rim.Price,
		rim.Qty,
		rim.ItemNum)
	db.Close()

	return
}

func UpdateRimQty(itemnum int, qty int) models.Rim {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateRimQty)
	util.CheckErr(err)

	statement.Exec(qty, itemnum)
	rim := GetRim(itemnum)
	db.Close()

	return rim
}
