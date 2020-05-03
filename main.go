package main

import (
	"bestprice_test/config"
	"bestprice_test/internal/app/repository"
	bp_sql "bestprice_test/internal/app/repository/sql"
	bp_http "bestprice_test/internal/http"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func createRouter(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()
	db, _ := bp_sql.ConnectToDB(cfg.MySQL)

	repo := repository.NewMapper(db)

	h := bp_http.Handler{Repo: repo, Authorization: cfg.JWT}

	r.HandleFunc("/categories/list", h.ListCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/{id}/read", h.ReadCategory).Methods(http.MethodGet)
	r.Handle("/categories/{id}/delete", h.Middleware(http.HandlerFunc(h.DeleteCategory))).Methods(http.MethodDelete)
	r.Handle("/categories/create", h.Middleware(http.HandlerFunc(h.CreateCategory))).Methods(http.MethodPost)
	r.Handle("/categories/update", h.Middleware(http.HandlerFunc(h.UpdateCategory))).Methods(http.MethodPut)

	r.HandleFunc("/products/list", h.ListProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}/read", h.ReadProduct).Methods(http.MethodGet)
	r.Handle("/products/{id}/delete", h.Middleware(http.HandlerFunc(h.DeleteProduct))).Methods(http.MethodDelete)
	r.Handle("/products/create", h.Middleware(http.HandlerFunc(h.CreateProduct))).Methods(http.MethodPost)
	r.Handle("/products/update", h.Middleware(http.HandlerFunc(h.UpdateProduct))).Methods(http.MethodPut)

	r.HandleFunc("/token", h.GetToken).Methods(http.MethodGet)

	return r

}

func main() {

	fmt.Println("App started.")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.NewConfig()
	r := createRouter(cfg)
	err = http.ListenAndServe(":5000", r)
	if err != nil {
		log.Panic(err)
	}

}
