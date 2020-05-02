package http

import (
	"bestprice_test/internal/app/repository"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {

	r := httptest.NewRequest("GET", "/products/list", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ListProducts)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadProduct(t *testing.T) {

	r := httptest.NewRequest("GET", "/products/{id}/read", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadProduct)
	expected := []byte(`{"success":true,"data":{"id":1,"category_id":0,"title":"","image_url":"","price":0,"description":"","created_at":"","updated_at":""}}`)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestReadProductIDNotInt(t *testing.T) {

	r := httptest.NewRequest("GET", "/product/{id}/read", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "a"})
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadProduct)
	expected := []byte(`{"success":false,"data":"Error: URL parameter 'a' is not an integer"}`)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}
func TestCreateProduct(t *testing.T) {

	jsonStr := []byte(`{
		"title":"bla",
		"category_id":1,
		"image_url":"blabla",
		"description":"Lorem ipsum dolor sit amet,",
		"price":19.99
	}`)
	r := httptest.NewRequest("POST", "/products/create", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	expected := []byte(`{"success":true,"data":null}`)

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateProduct)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestCreateProductWrongPayload(t *testing.T) {

	jsonStr := []byte(`{
		"category_id":1,
		"image_url":"blabla",
		"description":"Lorem ipsum dolor sit amet,",
		"price":19.99
	}`)
	r := httptest.NewRequest("POST", "/products/create", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	expected := []byte(`{"success":false,"data":{"title":["The title field is required!"]}}`)

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateProduct)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestUpdateProduct(t *testing.T) {

	r := httptest.NewRequest("PUT", "/products/update", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.UpdateProduct)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProduct(t *testing.T) {

	r := httptest.NewRequest("DELETE", "/products/1/delete", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.DeleteProduct)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListCategories(t *testing.T) {

	r := httptest.NewRequest("GET", "/categories/list", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ListCategories)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadCategory(t *testing.T) {

	r := httptest.NewRequest("GET", "/category/{id}/read", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadCategory)
	expected := []byte(`{"success":true,"data":{"id":1,"title":"","position":0,"image_url":"","created_at":"","updated_at":""}}`)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestReadCategoryIDNotInt(t *testing.T) {

	r := httptest.NewRequest("GET", "/category/{id}/read", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "a"})
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadCategory)
	expected := []byte(`{"success":false,"data":"Error: URL parameter 'a' is not an integer"}`)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestCreateCategory(t *testing.T) {

	jsonStr := []byte(`{
		"title":"Category Title",
		"position":1,
		"image_url":"image.com"
	}`)
	r := httptest.NewRequest("POST", "/categories/create", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	expected := []byte(`{"success":true,"data":null}`)

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateCategory)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestCreateCategoryWrongPayload(t *testing.T) {

	jsonStr := []byte(`{
		"position":1,
		"image_url":"image.com"
	}`)
	r := httptest.NewRequest("POST", "/categories/create", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	expected := []byte(`{"success":false,"data":{"title":["The title field is required!"]}}`)

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateCategory)

	srv.ServeHTTP(w, r)
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), strings.TrimRight(string(actual), "\n"))
}

func TestUpdateCategory(t *testing.T) {

	r := httptest.NewRequest("PUT", "/categories/update", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.UpdateCategory)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCategory(t *testing.T) {

	r := httptest.NewRequest("DELETE", "/categories/1/delete", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.DeleteCategory)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
