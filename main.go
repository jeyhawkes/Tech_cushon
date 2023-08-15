package main

import (
	"net/http"

	"github.com/jeyhawkes/tech_cushion/database"
	"github.com/jeyhawkes/tech_cushion/handlers"
	"github.com/jeyhawkes/tech_cushion/setup"
)

const (
	db_username = "root"
	db_password = "password"
	db_name     = "cushion"
)

func main() {

	var db database.Database

	if err := setup.Db(&db); err != nil {
		panic(err)
	}
	defer db.Close()

	investhttp := handlers.NewInvestmentHTTP(&db)
	http.HandleFunc("/invest/list/v1/", investhttp.HandleInvestment)
	http.HandleFunc("/invest/customer/v1/", investhttp.HandleCustomerInvestment)

	http.ListenAndServe(":8080", nil)
}
