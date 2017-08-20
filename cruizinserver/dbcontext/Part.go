package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetParts() ([]models.Part, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get parts")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetParts)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get parts")
	}
	defer rows.Close()
	var itemnum int
	var description string
	var condition string
	var price string
	parts := make([]models.Part, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &description, &condition, &price)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get parts")
		}
		parts = append(parts, models.Part{
			ItemNum:     itemnum,
			Description: description,
			Condition:   condition,
			Price:       price})
	}

	return parts, nil
}

func CreatePart(part models.Part) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add part")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreatePart)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add part")
	}
	defer statement.Close()
	_, err = statement.Exec(
		part.ItemNum,
		part.Description,
		part.Condition,
		part.Price)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add part")
	}

	return nil
}

func GetPart(key int) (models.Part, error) {
	part := models.Part{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return part, errors.New("Unable to get part")
	}
	defer db.Close()
	row, err := db.Query(queries.GetPart, key)
	if err != nil {
		log.Println(err)
		return part, errors.New("Unable to get part")
	}
	defer row.Close()
	var itemnum int
	var description string
	var condition string
	var price string
	if row.Next() {
		err = row.Scan(&itemnum, &description, &condition, &price)
		if err != nil {
			log.Println(err)
			return part, errors.New("Unable to get part")
		}
		part = models.Part{
			ItemNum:     itemnum,
			Description: description,
			Condition:   condition,
			Price:       price}
	}

	return part, nil
}

func DeletePart(itemnum int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete part")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeletePart)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete part")
	}
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete part")
	}

	return nil
}

func UpdatePart(part models.Part) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update part")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdatePart)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update part")
	}
	defer statement.Close()
	_, err = statement.Exec(
		part.Description,
		part.Condition,
		part.Price,
		part.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update part")
	}

	return nil
}
