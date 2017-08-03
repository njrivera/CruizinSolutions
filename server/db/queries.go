package db

import (
	"database/sql"

	"github.com/CruizinSolutions/server/utilities"
)

const (
	DBPATH                = "./TiosShop.sqlite"
	getAllCustomersClause = "SELECT * FROM Customers;"
	createCustomerClause  = "INSERT INTO Customers (firstname, lastname) VALUES (?, ?)"
)

type Customer struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func GetAllCustumers() []Customer {
	db, err := sql.Open("sqlite3", DBPATH)
	utilities.CheckErr(err)

	rows, err := db.Query(getAllCustomersClause)
	utilities.CheckErr(err)
	var id int
	var fname string
	var lname string
	var customers []Customer
	for rows.Next() {
		err = rows.Scan(&id, &fname, &lname)
		utilities.CheckErr(err)
		customers = append(customers, Customer{id, fname, lname})
	}
	rows.Close()
	db.Close()
	return customers
}

func CreateCustomer(cust Customer) {
	db, err := sql.Open("sqlite3", DBPATH)
	utilities.CheckErr(err)
	statement, _ := db.Prepare(createCustomerClause)
	statement.Exec(cust.Firstname, cust.Lastname)
	db.Close()
}
