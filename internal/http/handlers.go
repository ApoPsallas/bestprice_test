package http

import (
	"bestprice_test/internal/app/helper"
	"bestprice_test/internal/app/model"
	"bestprice_test/internal/app/repository"
	"bestprice_test/internal/app/services"
	"encoding/json"
	"io/ioutil"
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

	var p model.Product
	w.Header().Set("Content-type", "applciation/json")
	res := &helper.Response{Success: false, Data: nil}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: err.Error()}
		_ = json.NewEncoder(w).Encode(res)
		return
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: err.Error()}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	if validErrs := p.Validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: validErrs}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.SaveProduct(&p, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(res)
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

	var c model.Category
	w.Header().Set("Content-type", "applciation/json")
	res := &helper.Response{Success: false, Data: nil}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: err.Error()}
		_ = json.NewEncoder(w).Encode(res)
		return
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: err.Error()}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	if validErrs := c.Validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: validErrs}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.SaveCategory(&c, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(res)
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
