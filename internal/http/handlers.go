package http

import (
	"bestprice_test/config"
	"bestprice_test/internal/app/helper"
	"bestprice_test/internal/app/model"
	"bestprice_test/internal/app/repository"
	"bestprice_test/internal/app/services"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Handler struct {
	Authorization *config.JWTConfig
	Repo   *repository.Mapper
}

// PRODUCT HANDLERS

//ListProducts ...
func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := r.URL.Query()
	pn := helper.NewPagination(params)
	res := &helper.Response{Success: false, Data: nil}

	res = services.ListProducts(pn, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
	}
	_ = json.NewEncoder(w).Encode(res)
}

//ReadProduct ...
func (h *Handler) ReadProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	res := &helper.Response{Success: false, Data: nil}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: "Error: URL parameter '" + params["id"] + "' is not an integer"}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadProduct(id, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//CreateProduct ...
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var p model.Product
	w.Header().Set("Content-type", "application/json")
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

	if validErrs := validateProduct(&p); len(validErrs) > 0 {
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
	var p model.Product
	w.Header().Set("Content-type", "application/json")
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

	if validErrs := validateProduct(&p); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: validErrs}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadProduct(p.ID, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.UpdateProduct(&p, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//DeleteProduct ...
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	res := &helper.Response{Success: false, Data: nil}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: "Error: URL parameter '" + params["id"] + "' is not an integer"}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadProduct(id, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.DeleteProduct(id, h.Repo.ProductSqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//CATEGORY HANDLERS

//ListCategories ...
func (h *Handler) ListCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := r.URL.Query()
	pn := helper.NewPagination(params)
	res := &helper.Response{Success: false, Data: nil}

	res = services.ListCategories(pn, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//ReadCategory ...
func (h *Handler) ReadCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	res := &helper.Response{Success: false, Data: nil}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: "Error: URL parameter '" + params["id"] + "' is not an integer"}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadCategory(id, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//CreateCategory ...
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var c model.Category
	w.Header().Set("Content-type", "application/json")
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

	if validErrs := validateCategory(&c); len(validErrs) > 0 {
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
	var c model.Category
	w.Header().Set("Content-type", "application/json")
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

	if validErrs := validateCategory(&c); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: validErrs}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadCategory(c.ID, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.UpdateCategory(&c, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(res)
	}

	_ = json.NewEncoder(w).Encode(res)
}

//DeleteCategory ...
func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	res := &helper.Response{Success: false, Data: nil}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = &helper.Response{Success: false, Data: "Error: URL parameter '" + params["id"] + "' is not an integer"}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.ReadCategory(id, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res = services.DeleteCategory(id, h.Repo.CategorySqlMapper)
	if !res.Success {
		w.WriteHeader(http.StatusInternalServerError)

	}

	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "best_price"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(h.Authorization.JWTSecret))

	_ = json.NewEncoder(w).Encode(tokenString)
}

func (h *Handler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(h.Authorization.JWTSecret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func validateProduct(p *model.Product) url.Values {
	errs := url.Values{}

	if p.CategoryID == 0 {
		errs.Add("category_id", "The category_id field is required!")
	}

	if p.Title == "" {
		errs.Add("title", "The title field is required!")
	}

	if p.ImageURL == "" {
		errs.Add("image_url", "The image_url field is required!")
	}

	if p.Price == 0 {
		errs.Add("price", "The price field is required!")
	}

	if p.Description == "" {
		errs.Add("description", "The description field is required!")
	}

	return errs
}

func validateCategory(c *model.Category) url.Values {
	errs := url.Values{}

	if c.Title == "" {
		errs.Add("title", "The title field is required!")
	}

	if c.Position == 0 {
		errs.Add("position", "The position field is required!")
	}

	if c.ImageURL == "" {
		errs.Add("image_url", "The image_url field is required!")
	}

	return errs
}
