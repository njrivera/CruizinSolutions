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

	var partnum int
	var brand string
	var model string
	var size string
	var speedrating string
	var loadrange string
	var price float32
	var qty int
	var tires []models.Tire
	for rows.Next() {
		rows.Scan(&partnum, &brand, &model, &size, &speedrating, &loadrange, &price, &qty)
		tires = append(tires, models.Tire{
			PartNum:     partnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			SpeedRating: speedrating,
			LoadRange:   loadrange,
			Price:       price,
			Qty:         qty})
	}
	db.Close()

	return tires
}

func GetTire(key string) models.Tire {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetTire, key)
	util.CheckErr(err)

	var partnum int
	var brand string
	var model string
	var size string
	var speedrating string
	var loadrange string
	var price float32
	var qty int
	row.Scan(&partnum, &brand, &model, &size, &speedrating, &loadrange, &price, &qty)
	db.Close()
	tire := models.Tire{
		PartNum:     partnum,
		Brand:       brand,
		Model:       model,
		Size:        size,
		SpeedRating: speedrating,
		LoadRange:   loadrange,
		Price:       price,
		Qty:         qty}

	return tire
}

func CreateTire(tire models.Tire) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateTire)
	util.CheckErr(err)

	statement.Exec(
		tire.PartNum,
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.SpeedRating,
		tire.LoadRange,
		tire.Price,
		tire.Qty)
	db.Close()

	return
}
