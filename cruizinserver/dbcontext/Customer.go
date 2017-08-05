package dbcontext

import (
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/CruizinSolutions/cruizinserver/queries"
	"github.com/CruizinSolutions/cruizinserver/models"
)

func GetAllCustumers() []models.Customer {
	db, err := sql.Open("sqlite3", dbPath)
	util.CheckErr(err)
	rows, err := db.Query(queries.GetAllCustumers)
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
		customers = append(customers, Customer{id, fname, middle, lname, address, city, state, zip})
	}
	db.Close()
	return customers
}

func GetCustomer() models.Customer {
	db, err := sql.Open("sqlite3", dbPath)
	util.CheckErr(err)
	row, err := db.Query(queries.GetCustomer)
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
	return models.Customer{id, fname, middle, lname, address, city, state, zip}
}

func CreateCustomer(customer models.Customer) {
	db, err := sql.Open("sqlite3", dbPath)
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