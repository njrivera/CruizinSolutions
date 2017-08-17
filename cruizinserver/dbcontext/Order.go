package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func CreateOrder(order models.Order, items []models.ItemOrder) int {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateOrder)
	util.CheckErr(err)

	row, err := statement.Exec(
		order.Date,
		order.Cid,
		order.Vid,
		order.Odometer,
		order.Comments,
		order.Subtotal,
		order.Tax,
		order.Total)
	util.CheckErr(err)
	ordernum, err := row.LastInsertId()
	util.CheckErr(err)

	statement, err = db.Prepare(queries.CreateItemOrder)
	util.CheckErr(err)
	for _, item := range items {
		statement.Exec(
			ordernum,
			item.ItemNum,
			item.Price,
			item.Qty,
			item.Amount)
	}
	db.Close()

	return int(ordernum)
}

func GetOrders(key int) []models.OrderWithVehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetOrders, key)
	util.CheckErr(err)

	var ordernum int
	var date string
	var vid int
	var year int
	var make string
	var model string
	var odometer int
	var comments string
	var subtotal string
	var tax string
	var total string
	var orders []models.OrderWithVehicle
	for rows.Next() {
		rows.Scan(&ordernum, &date, &vid, &year, &make, &model, &odometer, &comments, &subtotal, &tax, &total)
		orders = append(orders, models.OrderWithVehicle{
			OrderNum: ordernum,
			Date:     date,
			Vid:      vid,
			Year:     year,
			Make:     make,
			Model:    model,
			Odometer: odometer,
			Comments: comments,
			Subtotal: subtotal,
			Tax:      tax,
			Total:    total})
	}
	db.Close()

	return orders
}

func GetItemOrders(key int) []models.ItemOrderWithDesc {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetItemOrders, key)
	util.CheckErr(err)

	var itemnum int
	var description string
	var price string
	var qty int
	var amount string
	var itemorders []models.ItemOrderWithDesc
	for rows.Next() {
		rows.Scan(&itemnum, &description, &price, &qty, &amount)
		itemorders = append(itemorders, models.ItemOrderWithDesc{
			ItemNum:     itemnum,
			Description: description,
			Price:       price,
			Qty:         qty,
			Amount:      amount})
	}
	db.Close()

	return itemorders
}

func GetCustVehicles(key int) []models.Vehicle {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetCustVehicles, key)
	util.CheckErr(err)

	var vid int
	var year int
	var make string
	var model string
	var vehicles []models.Vehicle
	for rows.Next() {
		rows.Scan(&vid, &year, &make, &model)
		vehicles = append(vehicles, models.Vehicle{
			Vid:   vid,
			Year:  year,
			Make:  make,
			Model: model})
	}
	db.Close()

	return vehicles
}
