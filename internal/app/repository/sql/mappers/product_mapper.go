package mappers

import (
	"bestprice_test/internal/app/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ProductSqlMapper struct {
	DB *sql.DB
}

func (m *ProductSqlMapper) Insert(p *model.Product) error {
	return nil
}

func (m *ProductSqlMapper) Read(ID int) (*model.Product, error) {
	return &model.Product{ID: ID}, nil
}

func (m *ProductSqlMapper) Update(p *model.Product) (*model.Product, error) {
	return p, nil
}

func (m *ProductSqlMapper) Delete(ID int) error {
	return nil
}

func (m *ProductSqlMapper) List() {

}
