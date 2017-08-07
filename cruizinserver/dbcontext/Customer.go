package dbcontext

import (
	"database/sql"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/util"
)

func GetCustomers() []models.Customer {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetCustomers)
	util.CheckErr(err)

	var cid int
	var name string
	var address string
	var city string
	var state string
	var zip string
	var phone string
	var customers []models.Customer
	for rows.Next() {
		rows.Scan(&cid, &name, &address, &city, &state, &zip, &phone)
		customers = append(customers, models.Customer{
			Cid:     cid,
			Name:    name,
			Address: address,
			City:    city,
			State:   state,
			Zipcode: zip,
			Phone:   phone})
	}
	db.Close()

	return customers
}

func CreateCustomer(customer models.Customer) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateCustomer)
	util.CheckErr(err)

	statement.Exec(
		customer.Name,
		customer.Address,
		customer.City,
		customer.State,
		customer.Zipcode,
		customer.Phone)
	db.Close()

	return
}

func GetCustomer(key int) models.Customer {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetCustomer, key)
	util.CheckErr(err)

	var cid int
	var name string
	var address string
	var city string
	var state string
	var zip string
	var phone string
	row.Scan(&cid, &name, &address, &city, &state, &zip, &phone)
	db.Close()
	customer := models.Customer{
		Name:    name,
		Address: address,
		City:    city,
		State:   state,
		Zipcode: zip,
		Phone:   phone}

	return customer
}

func DeleteCustomer(cid int) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.DeleteCustomer)
	util.CheckErr(err)

	statement.Exec(cid)

	return
}

func UpdateCustomer(customer models.Customer) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.UpdateCustomer)
	util.CheckErr(err)

	statement.Exec(
		customer.Name,
		customer.Address,
		customer.City,
		customer.State,
		customer.Zipcode,
		customer.Phone,
		customer.Cid)
	db.Close()

	return
}
