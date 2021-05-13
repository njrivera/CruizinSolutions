package queries

const (
	GetNewTireTax = "SELECT O.date, IO.qty, IO.amount FROM itemorders IO, items I, orders O WHERE I.type = \"NEW TIRE\" and IO.itemnum = I.itemnum and IO.ordernum = O.ordernum"
)
