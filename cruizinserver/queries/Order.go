package queries

const (
	CreateOrder     = "INSERT INTO orders (date, cid, vid, odometer, comments, subtotal, tax, total) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	GetCustVehicles = "SELECT DISTINCT V.vid, V.year, V.make, V.model FROM customers C, vehicles V, orders O WHERE C.cid = ? AND C.cid = O.cid AND O.vid = V.vid"
	CreateItemOrder = "INSERT INTO itemorders (ordernum, itemnum, qty, amount) VALUES (?, ?, ?, ?)"
)
