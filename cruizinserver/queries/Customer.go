package queries

const (
	GetCustomers   = "SELECT * FROM customers"
	GetCustomer    = "SELECT * FROM customers WHERE id = ?"
	CreateCustomer = "INSERT INTO customers (firstname, middle, lastname, address, city, state, zipcode) VALUES (?, ?, ?, ?, ?, ?, ?)"
)
