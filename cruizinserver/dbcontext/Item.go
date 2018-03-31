package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetItems() ([]models.Item, error) {
	db, err := sql.Open("sqlite3", GetDBPath())
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
	defer rows.Close()
	var itemnum int
	var description string
	var itype string
	items := make([]models.Item, 0)
	for rows.Next() {
		rows.Scan(&itemnum, &description, &itype)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get items")
		}
		items = append(items, models.Item{
			ItemNum:     itemnum,
			Description: description,
			Type:        itype})
	}

	return items, nil
}

func CreateItem(item models.Item) (int, error) {
	db, err := sql.Open("sqlite3", GetDBPath())
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
	defer statement.Close()
	row, err := statement.Exec(
		item.Description,
		item.Type)
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
	db, err := sql.Open("sqlite3", GetDBPath())
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
	defer row.Close()
	var itemnum int
	var description string
	var itype string
	if row.Next() {
		err = row.Scan(&itemnum, &description, &itype)
		if err != nil {
			log.Println(err)
			return item, errors.New("Unable to get item")
		}
		item = models.Item{
			ItemNum:     itemnum,
			Description: description,
			Type:        itype}
	}

	return item, nil
}

func DeleteItem(itemnum int) error {
	db, err := sql.Open("sqlite3", GetDBPath())
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
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete item")
	}

	return nil
}

func UpdateItem(item models.Item) error {
	db, err := sql.Open("sqlite3", GetDBPath())
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
	defer statement.Close()
	_, err = statement.Exec(
		item.Description,
		item.Type,
		item.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update item")
	}

	return nil
}
