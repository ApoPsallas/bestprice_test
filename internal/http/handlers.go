package http

import (
	"bestprice_test/internal/app/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service service.ApiService
}

//List ...
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode("Success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		_ = json.NewEncoder(w)
	}

}

//Read ...
func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode("Success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		_ = json.NewEncoder(w)
	}

}

//Create ...
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode("Success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		_ = json.NewEncoder(w)
	}
}

//Update ...
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode("Success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		_ = json.NewEncoder(w)
	}

}

//Delete ...
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode("Success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		_ = json.NewEncoder(w)
	}

}
