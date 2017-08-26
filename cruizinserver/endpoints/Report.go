package endpoints

import (
	"net/http"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"

	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterReportEndpoints(m *martini.ClassicMartini) {
	m.Group("/reports", func(r martini.Router) {
		r.Get("/newtiretax/:month/:year", getNewTireReportHandler)
	})
}

func getNewTireReportHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	report := models.NewTireTaxReport{}
	month := params["month"]
	year := params["year"]
	report, err := dbcontext.GetNewTireTax(month, year)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(report, w)
}
