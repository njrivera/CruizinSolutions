package queries

const (
	GetCustomers         = "SELECT * FROM customers"
	CreateCustomer       = "INSERT INTO customers (name, address, city, state, zipcode, phone, email) VALUES (?, ?, ?, ?, ?, ?, ?)"
	GetCustomer          = "SELECT * FROM customers WHERE cid = ?"
	DeleteCustomer       = "DELETE FROM customers WHERE cid = ?"
	UpdateCustomer       = "UPDATE customers SET name = ?, address = ?, city = ?, state = ?, zipcode = ?, phone = ?, email = ? WHERE cid = ?"
	GetSortedByOrderDate = "SELECT C.cid, C.name, C.address, C.city, C.state, C.zipcode, C.phone, C.email, O.date FROM customers C, orders O WHERE C.cid = O.cid"
)
