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

	var partnum int
	var name string
	var price float32
	var services []models.Service
	for rows.Next() {
		rows.Scan(&partnum, &name, &price)
		services = append(services, models.Service{
			PartNum: partnum,
			Name:    name,
			Price:   price})
	}
	db.Close()

	return services
}

func GetService(key string) models.Service {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetService, key)
	util.CheckErr(err)

	var partnum int
	var name string
	var price float32
	row.Scan(&partnum, &name, &price)
	db.Close()
	service := models.Service{
		PartNum: partnum,
		Name:    name,
		Price:   price}

	return service
}

func CreateService(service models.Service) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateService)
	util.CheckErr(err)

	statement.Exec(
		service.PartNum,
		service.Name,
		service.Price)
	db.Close()

	return
}
