package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func CreateOrder(order models.Order, items []models.ItemOrder) (int, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to create order")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateOrder)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to create order")
	}

	row, err := statement.Exec(
		order.Date,
		order.Cid,
		order.Vid,
		order.Odometer,
		order.Comments,
		order.Subtotal,
		order.Tax,
		order.Total)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to create order")
	}
	ordernum, err := row.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to get ordernum")
	}

	statement, err = db.Prepare(queries.CreateItemOrder)
	if err != nil {
		log.Println(err)
		return -1, errors.New("Unable to create item orders")
	}
	for _, item := range items {
		_, err = statement.Exec(
			ordernum,
			item.ItemNum,
			item.Price,
			item.Qty,
			item.Amount)
		if err != nil {
			log.Println(err)
			return -1, errors.New("Unable to create item orders")
		}
	}

	return int(ordernum), nil
}

func GetOrders(key int) ([]models.OrderWithVehicle, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get orders")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetOrders, key)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get orders")
	}

	var ordernum int
	var date string
	var vid int
	var year string
	var vmake string
	var model string
	var odometer int
	var comments string
	var subtotal string
	var tax string
	var total string
	orders := make([]models.OrderWithVehicle, 0)
	for rows.Next() {
		err = rows.Scan(&ordernum, &date, &vid, &year, &vmake, &model, &odometer, &comments, &subtotal, &tax, &total)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get orders")
		}
		orders = append(orders, models.OrderWithVehicle{
			OrderNum: ordernum,
			Date:     date,
			Vid:      vid,
			Year:     year,
			Make:     vmake,
			Model:    model,
			Odometer: odometer,
			Comments: comments,
			Subtotal: subtotal,
			Tax:      tax,
			Total:    total})
	}
	return orders, nil
}

func GetItemOrders(key int) ([]models.ItemOrderWithDesc, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get item orders")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetItemOrders, key)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get item orders")
	}

	var itemnum int
	var description string
	var price string
	var qty int
	var amount string
	itemorders := make([]models.ItemOrderWithDesc, 0)
	for rows.Next() {
		err = rows.Scan(&itemnum, &description, &price, &qty, &amount)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get item orders")
		}
		itemorders = append(itemorders, models.ItemOrderWithDesc{
			ItemNum:     itemnum,
			Description: description,
			Price:       price,
			Qty:         qty,
			Amount:      amount})
	}

	return itemorders, nil
}

func GetCustVehicles(key int) ([]models.Vehicle, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get customer's vehicles")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetCustVehicles, key)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get customer's vehicles")
	}

	var vid int
	var year string
	var vmake string
	var model string
	vehicles := make([]models.Vehicle, 0)
	for rows.Next() {
		err = rows.Scan(&vid, &year, &vmake, &model)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get customer's vehicles")
		}
		vehicles = append(vehicles, models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  vmake,
			Model: model})
	}

	return vehicles, nil
}
