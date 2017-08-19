package dbcontext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CruizinSolutions/cruizinserver/database"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetCustomers() ([]models.Customer, error) {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get customers")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetCustomers)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to get customers")
	}

	var cid int
	var name string
	var address string
	var city string
	var state string
	var zip string
	var phone string
	customers := make([]models.Customer, 0)
	for rows.Next() {
		err = rows.Scan(&cid, &name, &address, &city, &state, &zip, &phone)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Unable to get customers")
		}
		customers = append(customers, models.Customer{
			Cid:     cid,
			Name:    name,
			Address: address,
			City:    city,
			State:   state,
			Zipcode: zip,
			Phone:   phone})
	}

	return customers, nil
}

func CreateCustomer(customer models.Customer) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add customer")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.CreateCustomer)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add customer")
	}

	_, err = statement.Exec(
		customer.Name,
		customer.Address,
		customer.City,
		customer.State,
		customer.Zipcode,
		customer.Phone)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to add customer")
	}

	return nil
}

func GetCustomer(key int) (models.Customer, error) {
	customer := models.Customer{}
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return customer, errors.New("Unable to get customer")
	}
	defer db.Close()
	row, err := db.Query(queries.GetCustomer, key)
	if err != nil {
		log.Println(err)
		return customer, errors.New("Unable to get customer")
	}

	var cid int
	var name string
	var address string
	var city string
	var state string
	var zip string
	var phone string
	err = row.Scan(&cid, &name, &address, &city, &state, &zip, &phone)
	if err != nil {
		log.Println(err)
		return customer, errors.New("Unable to get customer")
	}
	customer = models.Customer{
		Name:    name,
		Address: address,
		City:    city,
		State:   state,
		Zipcode: zip,
		Phone:   phone}

	return customer, nil
}

func DeleteCustomer(cid int) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete customer")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.DeleteCustomer)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete customer")
	}

	_, err = statement.Exec(cid)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to delete customer")
	}

	return nil
}

func UpdateCustomer(customer models.Customer) error {
	db, err := sql.Open("sqlite3", database.DBPath)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update customer")
	}
	defer db.Close()
	statement, err := db.Prepare(queries.UpdateCustomer)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update customer")
	}

	_, err = statement.Exec(
		customer.Name,
		customer.Address,
		customer.City,
		customer.State,
		customer.Zipcode,
		customer.Phone,
		customer.Cid)
	if err != nil {
		log.Println(err)
		return errors.New("Unable to update customer")
	}

	return nil
}
