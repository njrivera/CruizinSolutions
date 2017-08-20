package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetServices() ([]models.Service, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get services")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetServices)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get services")
	}
	defer rows.Close()
	var itemnum int
	var description string
	var price string
	services := make([]models.Service, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &description, &price)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get services")
		}
		services = append(services, models.Service{
			ItemNum:     itemnum,
			Description: description,
			Price:       price})
	}

	return services, nil
}

func CreateService(service models.Service) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add service")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateService)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add service")
	}
	defer statement.Close()
	_, err = statement.Exec(
		service.ItemNum,
		service.Description,
		service.Price)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add service")
	}

	return nil
}

func GetService(key int) (models.Service, error) {
	service := models.Service{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return service, errors.New("Unable to get service")
	}
	defer db.Close()
	row, err := db.Query(queries.GetService, key)
	if err != nil {
		log.Println(err)
		return service, errors.New("Unable to get service")
	}
	defer row.Close()
	var itemnum int
	var description string
	var price string
	if row.Next() {
		err = row.Scan(&itemnum, &description, &price)
		if err != nil {
			log.Println(err)
			return service, errors.New("Unable to get service")
		}
		service = models.Service{
			ItemNum:     itemnum,
			Description: description,
			Price:       price}
	}

	return service, nil
}

func DeleteService(itemnum int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete service")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteService)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete service")
	}
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete service")
	}

	return nil
}

func UpdateService(service models.Service) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update service")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateService)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update service")
	}
	defer statement.Close()
	_, err = statement.Exec(
		service.Description,
		service.Price,
		service.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update service")
	}

	return nil
}
