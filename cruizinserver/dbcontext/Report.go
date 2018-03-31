package dbcontext

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/CruizinSolutions/cruizinserver/models"

	"github.com/CruizinSolutions/cruizinserver/queries"
)

func GetNewTireTax(month string, year string) (models.NewTireTaxReport, error) {
	report := models.NewTireTaxReport{}
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return report, errors.New("Unable to get report")
	}
	defer db.Close()
	rows, err := db.Query(queries.GetNewTireTax)
	if err != nil {
		log.Println(err)
		return report, errors.New("Unable to get report")
	}
	defer rows.Close()

	var date string
	var qty int
	var amount string
	tires := make([]models.NewTire, 0)
	for rows.Next() {
		err = rows.Scan(&date, &qty, &amount)
		if err != nil {
			log.Println(err)
			return report, errors.New("Unable to get report")
		}
		tires = append(tires, models.NewTire{
			Date:   date,
			Qty:    qty,
			Amount: amount})
	}

	var d []string
	var reportTax float64
	var tireTax float64
	for _, tire := range tires {
		d = strings.Split(tire.Date, "/")
		if d[2] == year && d[0] == month {
			report.Qty += tire.Qty
			reportTax, _ = strconv.ParseFloat(report.Tax, 64)
			tireTax, _ = strconv.ParseFloat(tire.Amount, 64)
			report.Tax = roundCents(reportTax + tireTax*.07)
		}
	}

	return report, nil
}

func roundCents(num float64) string {
	front := int(num)
	back := int(num*1000) - front*1000
	if int(back)%10 > 4 {
		back += 10
	}
	return strconv.Itoa(front) + "." + strconv.Itoa(int(back/10))
}
