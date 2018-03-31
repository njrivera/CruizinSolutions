package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetTires() ([]models.Tire, error) {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get tires")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetTires)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get tires")
	}
	defer rows.Close()
	var itemnum int
	var brand string
	var model string
	var size string
	var servicedesc string
	var warranty string
	var condition string
	var price string
	var qty int
	tires := make([]models.Tire, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &brand, &model, &size, &servicedesc, &warranty, &condition, &price, &qty)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get tires")
		}
		tires = append(tires, models.Tire{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			ServiceDesc: servicedesc,
			Warranty:    warranty,
			Condition:   condition,
			Price:       price,
			Qty:         qty})
	}

	return tires, nil
}

func CreateTire(tire models.Tire) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add tire")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateTire)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add tire")
	}
	defer statement.Close()
	_, err = statement.Exec(
		tire.ItemNum,
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.ServiceDesc,
		tire.Warranty,
		tire.Condition,
		tire.Price,
		tire.Qty)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add tire")
	}

	return nil
}

func GetTire(key int) (models.Tire, error) {
	tire := models.Tire{}
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to get tire")
	}
	defer db.Close()
	row, err := db.Query(queries.GetTire, key)
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to get tire")
	}
	defer row.Close()
	var itemnum int
	var brand string
	var model string
	var size string
	var servicedesc string
	var warranty string
	var condition string
	var price string
	var qty int
	if row.Next() {
		err = row.Scan(&itemnum, &brand, &model, &size, &servicedesc, &warranty, &condition, &price, &qty)
		if err != nil {
			log.Println(err)
			return tire, errors.New("Unable to get tire")
		}
		tire = models.Tire{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			ServiceDesc: servicedesc,
			Warranty:    warranty,
			Condition:   condition,
			Price:       price,
			Qty:         qty}
	}

	return tire, nil
}

func DeleteTire(itemnum int) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete tire")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteTire)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete tire")
	}
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete tire")
	}

	return nil
}

func UpdateTire(tire models.Tire) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update tire")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateTire)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update tire")
	}
	defer statement.Close()
	_, err = statement.Exec(
		tire.Brand,
		tire.Model,
		tire.Size,
		tire.ServiceDesc,
		tire.Warranty,
		tire.Condition,
		tire.Price,
		tire.Qty,
		tire.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update tire")
	}

	return nil
}

func UpdateTireQty(itemnum int, qty int) (models.Tire, error) {
	tire := models.Tire{}
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to update tire qty")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateTireQty)
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to update tire qty")
	}
	defer statement.Close()
	_, err = statement.Exec(qty, itemnum)
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to update tire qty")
	}
	tire, err = GetTire(itemnum)
	if err != nil {
		log.Println(err)
		return tire, errors.New("Unable to get tire")
	}

	return tire, nil
}
