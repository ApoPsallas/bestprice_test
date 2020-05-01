package repository

import (
	"bestprice_test/internal/app/repository/sql/mappers"
	"database/sql"
)

type Mapper struct {
	Cm *mappers.CategorySqlMapper
	Pm *mappers.ProductSqlMapper
}

func NewMapper(DB *sql.DB) *Mapper {
	cm := &mappers.CategorySqlMapper{DB: DB}
	pm := &mappers.ProductSqlMapper{DB: DB}
	return &Mapper{Cm: cm, Pm: pm}

}
