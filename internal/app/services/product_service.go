package services

import (
	"bestprice_test/internal/app/model"
	"bestprice_test/internal/app/repository/sql/mappers"

	"bestprice_test/internal/app/helper"
)

func SaveProduct(c *model.Product, m *mappers.ProductSqlMapper) *helper.Response {
	err := m.Insert(c)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: nil}
}

func ReadProduct(ID int, m *mappers.ProductSqlMapper) *helper.Response {
	p, err := m.Read(ID)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: p}
}

func UpdateProduct(c *model.Product, m *mappers.ProductSqlMapper) *helper.Response {
	err := m.Update(c)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: nil}
}

func DeleteProduct(ID int, m *mappers.ProductSqlMapper) *helper.Response {
	err := m.Delete(ID)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: nil}
}

func ListProducts(pn *helper.Pagination, m *mappers.ProductSqlMapper) *helper.Response {
	p, err := m.List(pn)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: p}
}
