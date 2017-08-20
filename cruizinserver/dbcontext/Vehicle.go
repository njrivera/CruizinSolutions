package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetVehicles() ([]models.Vehicle, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get vehicles")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetVehicles)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get vehicles")
	}
	defer rows.Close()
	var vid int
	var year string
	var vmake string
	var model string
	vehicles := make([]models.Vehicle, 0)
	for rows.Next() {
		err = rows.Scan(&vid, &year, &vmake, &model)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get vehicles")
		}
		vehicles = append(vehicles, models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  vmake,
			Model: model})
	}

	return vehicles, nil
}

func CreateVehicle(vehicle models.Vehicle) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add vehicle")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateVehicle)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add vehicle")
	}
	defer statement.Close()
	_, err = statement.Exec(
		vehicle.Year,
		vehicle.Make,
		vehicle.Model)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add vehicle")
	}

	return nil
}

func GetVehicle(key int) (models.Vehicle, error) {
	vehicle := models.Vehicle{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return vehicle, errors.New("Unable to get vehicle")
	}
	defer db.Close()
	row, err := db.Query(queries.GetVehicle, key)
	if err != nil {
		log.Println(err)
		return vehicle, errors.New("Unable to get vehicle")
	}
	defer row.Close()
	var vid int
	var year string
	var make string
	var model string
	if row.Next() {
		err = row.Scan(&vid, &year, &make, &model)
		if err != nil {
			log.Println(err)
			return vehicle, errors.New("Unable to get vehicle")
		}
		vehicle = models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  make,
			Model: model}
	}

	return vehicle, nil
}

func DeleteVehicle(vid int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete vehicle")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteVehicle)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete vehicle")
	}
	defer statement.Close()
	_, err = statement.Exec(vid)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete vehicle")
	}

	return nil
}

func UpdateVehicle(vehicle models.Vehicle) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update vehicle")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateVehicle)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update vehicle")
	}
	defer statement.Close()
	_, err = statement.Exec(
		vehicle.Year,
		vehicle.Make,
		vehicle.Model,
		vehicle.Vid)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update vehicle")
	}

	return nil
}
