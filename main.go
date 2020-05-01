package main

import (
	"bestprice_test/internal/app/repository"
	app_sql "bestprice_test/internal/app/repository/sql"
	bp_http "bestprice_test/internal/http"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createRouter() *http.ServeMux {
	r := http.NewServeMux()
	db, _ := app_sql.ConnectToDB()

	repo := repository.NewMapper(db)

	h := bp_http.Handler{Repo: repo}

	r.HandleFunc("/list", h.List)
	r.HandleFunc("/read", h.Read)
	r.HandleFunc("/create", h.Create)
	r.HandleFunc("/update", h.Update)
	r.HandleFunc("/delete", h.Delete)
	return r

}

func main() {

	fmt.Println("App started.")
	r := createRouter()
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Panic(err)
	}

}
