package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/server/models"
	"github.com/CruizinSolutions/server/queries"
)

func GetRims() ([]models.Rim, error) {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get rims")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetRims)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get rims")
	}
	defer rows.Close()
	var itemnum int
	var brand string
	var model string
	var size string
	var boltpattern string
	var finish string
	var condition string
	var price string
	var qty int
	rims := make([]models.Rim, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &brand, &model, &size, &boltpattern, &finish, &condition, &price, &qty)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get rims")
		}
		rims = append(rims, models.Rim{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			BoltPattern: boltpattern,
			Finish:      finish,
			Condition:   condition,
			Price:       price,
			Qty:         qty})
	}

	return rims, nil
}

func CreateRim(rim models.Rim) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add rim")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateRim)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add rim")
	}
	defer statement.Close()
	_, err = statement.Exec(
		rim.ItemNum,
		rim.Brand,
		rim.Model,
		rim.Size,
		rim.BoltPattern,
		rim.Finish,
		rim.Condition,
		rim.Price,
		rim.Qty)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add rim")
	}

	return nil
}

func GetRim(key int) (models.Rim, error) {
	rim := models.Rim{}
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to get rim")
	}
	defer db.Close()
	row, err := db.Query(queries.GetRim, key)
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to get rim")
	}
	defer row.Close()
	var itemnum int
	var brand string
	var model string
	var size string
	var boltpattern string
	var finish string
	var condition string
	var price string
	var qty int
	if row.Next() {
		err = row.Scan(&itemnum, &brand, &model, &size, &boltpattern, &finish, &condition, &price, &qty)
		if err != nil {
			log.Println(err)
			return rim, errors.New("Unable to get rim")
		}
		rim = models.Rim{
			ItemNum:     itemnum,
			Brand:       brand,
			Model:       model,
			Size:        size,
			BoltPattern: boltpattern,
			Finish:      finish,
			Condition:   condition,
			Price:       price,
			Qty:         qty}
	}

	return rim, nil
}

func DeleteRim(itemnum int) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete rim")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteRim)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete rim")
	}
	defer statement.Close()
	_, err = statement.Exec(itemnum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete rim")
	}

	return nil
}

func UpdateRim(rim models.Rim) error {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update rim")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateRim)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update rim")
	}
	defer statement.Close()
	_, err = statement.Exec(
		rim.Brand,
		rim.Model,
		rim.Size,
		rim.BoltPattern,
		rim.Finish,
		rim.Condition,
		rim.Price,
		rim.Qty,
		rim.ItemNum)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update rim")
	}

	return nil
}

func UpdateRimQty(itemnum int, qty int) (models.Rim, error) {
	rim := models.Rim{}
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to update rim qty")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateRimQty)
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to update rim qty")
	}
	defer statement.Close()
	_, err = statement.Exec(qty, itemnum)
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to update rim qty")
	}
	rim, err = GetRim(itemnum)
	if err != nil {
		log.Println(err)
		return rim, errors.New("Unable to get rim")
	}

	return rim, nil
}
