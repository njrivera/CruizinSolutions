package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetServices() []models.Service {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetServices)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price float32
	var services []models.Service
	for rows.Next() {
		rows.Scan(&itemnum, &description, &price)
		services = append(services, models.Service{
			ItemNum:     itemnum,
			Description: description,
			Price:       price})
	}
	db.Close()

	return services
}

func CreateService(service models.Service) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateService)
	util.CheckErr(err)

	statement.Exec(
		service.ItemNum,
		service.Description,
		service.Price)
	db.Close()

	return
}

func GetService(key int) models.Service {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetService, key)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price float32
	row.Scan(&itemnum, &description, &price)
	db.Close()
	service := models.Service{
		ItemNum:     itemnum,
		Description: description,
		Price:       price}

	return service
}

func DeleteService(itemnum int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteService)
	util.CheckErr(err)

	statement.Exec(itemnum)

	return
}

func UpdateService(service models.Service) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateService)
	util.CheckErr(err)

	statement.Exec(
		service.Description,
		service.Price,
		service.ItemNum)
	db.Close()

	return
}
