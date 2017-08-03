package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

type customer struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	m := martini.Classic()
	m.Get("/", func() (int, string) {
		return http.StatusOK, "You're a cunt!"
	})
	m.Get("/names", getNamesHandler)
	m.Post("/create", createCustomer)
	m.Run()
}

func getNamesHandler(r *http.Request, w http.ResponseWriter) int {
	db, err := sql.Open("sqlite3", "./TiosShop.sqlite")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM Customers;")
	checkErr(err)
	var id int
	var fname string
	var lname string
	var names []customer
	for rows.Next() {
		err = rows.Scan(&id, &fname, &lname)
		checkErr(err)
		names = append(names, customer{id, fname, lname})
	}
	rows.Close()
	db.Close()

	js, _ := json.MarshalIndent(names, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return http.StatusOK
}

func createCustomer(r *http.Request) int {
	customer := customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	checkErr(err)

	db, err := sql.Open("sqlite3", "./TiosShop.sqlite")
	statement, _ := db.Prepare("INSERT INTO Customers (firstname, lastname) VALUES (?, ?)")
	statement.Exec(customer.Firstname, customer.Lastname)
	db.Close()
	return http.StatusOK
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
