package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetPackages() ([]models.Package, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get packages")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetPackages)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get packages")
	}
	defer rows.Close()
	var itemnum int
	var description string
	var price string
	packages := make([]models.Package, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &description, &price)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get packages")
		}
		packages = append(packages, models.Package{
			ItemNum:     itemnum,
			Description: description,
			Price:       price})
	}

	return packages, nil
}

func CreatePackage(pack models.Package) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add package")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreatePackage)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add package")
	}
	defer statement.Close()
	_, err = statement.Exec(
		pack.ItemNum,
		pack.Description,
		pack.Price)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add package")
	}

	return nil
}

func GetPackage(key int) (models.Package, error) {
	pack := models.Package{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return pack, errors.New("Unable to get package")
	}
	defer db.Close()
	row, err := db.Query(queries.GetPackage, key)
	if err != nil {
		log.Println(err)
		return pack, errors.New("Unable to get package")
	}
	defer row.Close()
	var itemnum int
	var description string
	var price string
	if row.Next() {
		err = row.Scan(&itemnum, &description, &price)
		if err != nil {
			log.Println(err)
			return pack, errors.New("Unable to get package")
		}
		pack = models.Package{
			ItemNum:     itemnum,
			Description: description,
			Price:       price}
	}

	return pack, nil
}

func DeletePackage(itemnum int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete package")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeletePackage)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete package")
	}
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete package")
	}

	return nil
}

func UpdatePackage(pack models.Package) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update package")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdatePackage)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update package")
	}
	defer statement.Close()
	_, err = statement.Exec(
		pack.Description,
		pack.Price,
		pack.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update package")
	}

	return nil
}
