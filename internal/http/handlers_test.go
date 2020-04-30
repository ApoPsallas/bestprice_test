package http

import (
	"bestprice_test/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	request := httptest.NewRequest("GET", "/list", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.List)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

func TestListWrongMethod(t *testing.T) {

	request := httptest.NewRequest("POST", "/list", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.List)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusMethodNotAllowed, responseWriter.Code)
}

func TestRead(t *testing.T) {

	request := httptest.NewRequest("GET", "/read", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Read)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

//http.StatusMethodNotAllowed

func TestReadWrongMethod(t *testing.T) {

	request := httptest.NewRequest("POST", "/read", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Read)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusMethodNotAllowed, responseWriter.Code)
}
func TestCreate(t *testing.T) {

	request := httptest.NewRequest("POST", "/create", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Create)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

//http.StatusMethodNotAllowed

func TestCreateWrongMethod(t *testing.T) {

	request := httptest.NewRequest("GET", "/create", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Create)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusMethodNotAllowed, responseWriter.Code)
}

func TestUpdate(t *testing.T) {

	request := httptest.NewRequest("PUT", "/update", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Update)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

//http.StatusMethodNotAllowed

func TestUpdateWrongMethod(t *testing.T) {

	request := httptest.NewRequest("POST", "/update", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Update)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusMethodNotAllowed, responseWriter.Code)
}

func TestDelete(t *testing.T) {

	request := httptest.NewRequest("DELETE", "/delete", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Delete)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

//http.StatusMethodNotAllowed

func TestDeleteWrongMethod(t *testing.T) {

	request := httptest.NewRequest("POST", "/delete", nil)
	responseWriter := httptest.NewRecorder()

	s := service.ApiService{}

	h := Handler{Service: s}

	srv := http.HandlerFunc(h.Delete)

	srv.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusMethodNotAllowed, responseWriter.Code)
}
