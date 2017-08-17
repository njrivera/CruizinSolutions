package queries

const (
	CreateOrder     = "INSERT INTO orders (date, cid, vid, odometer, comments, subtotal, tax, total) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	GetCustVehicles = "SELECT DISTINCT V.vid, V.year, V.make, V.model FROM customers C, vehicles V, orders O WHERE C.cid = ? AND C.cid = O.cid AND O.vid = V.vid"
	CreateItemOrder = "INSERT INTO itemorders (ordernum, itemnum, price, qty, amount) VALUES (?, ?, ?, ?, ?)"
	GetOrders       = "SELECT O.ordernum, O.date, V.vid, V.year, V.make, V.model, O.odometer, O.comments, O.subtotal, O.tax, O.total FROM orders O, vehicles V WHERE cid = ? AND V.vid = O.vid"
	GetItemOrders   = "SELECT I.itemnum, I.description, IO.price, IO.qty, IO.amount FROM itemorders IO, items I WHERE IO.ordernum = ? AND I.itemnum = IO.itemnum"
)
