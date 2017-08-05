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

	var id int
	var fname string
	var middle string
	var lname string
	var address string
	var city string
	var state string
	var zip string
	var customers []models.Customer
	for rows.Next() {
		rows.Scan(&id, &fname, &middle, &lname, &address, &city, &state, &zip)
		customers = append(customers, models.Customer{
			Id:        id,
			Firstname: fname,
			Middle:    middle,
			Lastname:  lname,
			Address:   address,
			City:      city,
			State:     state,
			Zipcode:   zip})
	}
	db.Close()

	return customers
}

func GetCustomer(key string) models.Customer {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetCustomer, key)
	util.CheckErr(err)

	var id int
	var fname string
	var middle string
	var lname string
	var address string
	var city string
	var state string
	var zip string
	row.Scan(&id, &fname, &middle, &lname, &address, &city, &state, &zip)
	db.Close()
	customer := models.Customer{
		Id:        id,
		Firstname: fname,
		Middle:    middle,
		Lastname:  lname,
		Address:   address,
		City:      city,
		State:     state,
		Zipcode:   zip}

	return customer
}

func CreateCustomer(customer models.Customer) {
	db, err := sql.Open("sqlite3", database.DBPath)
	util.CheckErr(err)
	statement, err := db.Prepare(queries.CreateCustomer)
	util.CheckErr(err)

	statement.Exec(
		customer.Firstname,
		customer.Middle,
		customer.Lastname,
		customer.Address,
		customer.City,
		customer.State,
		customer.Zipcode)
	db.Close()

	return
}
