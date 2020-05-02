package services

import (
	"bestprice_test/internal/app/model"
	"bestprice_test/internal/app/repository/sql/mappers"

	"bestprice_test/internal/app/helper"
)

func SaveCategory(c *model.Category, m *mappers.CategorySqlMapper) *helper.Response {
	err := m.Insert(c)
	if err != nil {
		return &helper.Response{Success: false, Data: err}
	}
	return &helper.Response{Success: true, Data: nil}
}

func ReadCategory(ID int, m *mappers.CategorySqlMapper) *helper.Response {
	p, err := m.Read(ID)
	if err != nil {
		return &helper.Response{Success: false, Data: err.Error}
	}
	return &helper.Response{Success: true, Data: p}
}
