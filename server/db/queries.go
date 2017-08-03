package db

import (
	"database/sql"

	"github.com/CruizinSolutions/server/utilities"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath                = "db/TiosShop.sqlite"
	getAllCustomersClause = "SELECT * FROM Customers;"
	createCustomerClause  = "INSERT INTO Customers (firstname, lastname) VALUES (?, ?)"
)

type Customer struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func GetAllCustumers() []Customer {
	db, err := sql.Open("sqlite3", dbPath)
	utilities.CheckErr(err)
	rows, err := db.Query(getAllCustomersClause)
	utilities.CheckErr(err)
	var id int
	var fname string
	var lname string
	var customers []Customer
	for rows.Next() {
		rows.Scan(&id, &fname, &lname)
		customers = append(customers, Customer{id, fname, lname})
	}
	db.Close()
	return customers
}

func CreateCustomer(customer Customer) {
	db, err := sql.Open("sqlite3", dbPath)
	utilities.CheckErr(err)
	statement, err := db.Prepare(createCustomerClause)
	utilities.CheckErr(err)
	statement.Exec(customer.Firstname, customer.Lastname)
	db.Close()
	return
}
