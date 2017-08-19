package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetItems() ([]models.Item, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get items")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetItems)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get items")
	}

	var itemnum int
	var description string
	items := make([]models.Item, 0)
	for rows.Next() {
		rows.Scan(&itemnum, &description)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get items")
		}
		items = append(items, models.Item{
			ItemNum:     itemnum,
			Description: description})
	}

	return items, nil
}

func CreateItem(item models.Item) (int, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to add item")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateItem)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to add item")
	}

	row, err := statement.Exec(
		item.Description)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to add item")
	}
	id, err := row.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to add item")
	}

	return int(id), nil
}

func GetItem(key int) (models.Item, error) {
	item := models.Item{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return item, errors.New("Unable to get item")
	}
	defer db.Close()
	row, err := db.Query(queries.GetItem, key)
	if err != nil {
		log.Println(err)
		return item, errors.New("Unable to get item")
	}

	var itemnum int
	var description string
	err = row.Scan(&itemnum, &description)
	if err != nil {
		log.Println(err)
		return item, errors.New("Unable to get item")
	}
	item = models.Item{
		ItemNum:     itemnum,
		Description: description}

	return item, nil
}

func DeleteItem(itemnum int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete item")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteItem)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete item")
	}
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete item")
	}

	return nil
}

func UpdateItem(itemnum int, description string) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update item")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateItem)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update item")
	}

	_, err = statement.Exec(
		description,
		itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update item")
	}

	return nil
}
