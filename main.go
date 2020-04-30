package main

import (
	"bestprice_test/internal/app/service"
	bp_http "bestprice_test/internal/http"
	"fmt"
	"log"
	"net/http"
)

func createRouter() *http.ServeMux {
	r := http.NewServeMux()
	s := service.ApiService{}
	h := bp_http.Handler{Service: s}
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
