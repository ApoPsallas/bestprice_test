package http

import (
	"bestprice_test/internal/app/repository"
	"net/http"
	"net/http/httptest"
	"testing"

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

	r := httptest.NewRequest("GET", "/products/1/read", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadProduct)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateProduct(t *testing.T) {

	r := httptest.NewRequest("POST", "/products/create", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateProduct)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
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

	r := httptest.NewRequest("GET", "/categories/1/read", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.ReadCategory)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateCategory(t *testing.T) {

	r := httptest.NewRequest("POST", "/categories/create", nil)
	w := httptest.NewRecorder()

	repo := repository.Mapper{}
	h := Handler{Repo: &repo}

	srv := http.HandlerFunc(h.CreateCategory)

	srv.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
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
