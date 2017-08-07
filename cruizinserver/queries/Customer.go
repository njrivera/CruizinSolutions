package queries

const (
	GetCustomers   = "SELECT * FROM customers"
	CreateCustomer = "INSERT INTO customers (name, address, city, state, zipcode, phone) VALUES (?, ?, ?, ?, ?, ?)"
	GetCustomer    = "SELECT * FROM customers WHERE cid = ?"
	DeleteCustomer = "DELETE FROM customers WHERE cid = ?"
	UpdateCustomer = "UPDATE customers SET name = ?, address = ?, city = ?, state = ?, zipcode = ?, phone = ? WHERE cid = ?"
)
