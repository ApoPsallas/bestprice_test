package repository

import (
	"bestprice_test/internal/app/repository/sql/mappers"
	"database/sql"
)

type Mapper struct {
	CategorySqlMapper *mappers.CategorySqlMapper
	ProductSqlMapper  *mappers.ProductSqlMapper
}

func NewMapper(DB *sql.DB) *Mapper {
	cm := &mappers.CategorySqlMapper{DB: DB}
	pm := &mappers.ProductSqlMapper{DB: DB}
	return &Mapper{CategorySqlMapper: cm, ProductSqlMapper: pm}

}
