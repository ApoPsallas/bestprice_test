package http

import (
	"bestprice_test/internal/app/repository"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repo *repository.Mapper
}

// PRODUCT HANDLERS

//ListProducts ...
func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//ReadProduct ...
func (h *Handler) ReadProduct(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//CreateProduct ...
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//UpdateProduct ...
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//DeleteProduct ...
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//CATEGORY HANDLERS

//ListCategories ...
func (h *Handler) ListCategories(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//ReadCategory ...
func (h *Handler) ReadCategory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//CreateCategory ...
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//UpdateCategory ...
func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}

//DeleteCategory ...
func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Success")
}
