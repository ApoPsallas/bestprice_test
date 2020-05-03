package mappers

import (
	"bestprice_test/internal/app/helper"
	"bestprice_test/internal/app/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ProductSqlMapper struct {
	DB *sql.DB
}

func (m *ProductSqlMapper) Insert(p *model.Product) error {
	q := "INSERT INTO products (category_id,title,image_url,price,description) VALUES (?,?,?,?,?);"
	_, err := m.DB.Query(q, p.CategoryID, p.Title, p.ImageURL, p.Price, p.Description)
	return err
}

func (m *ProductSqlMapper) Read(ID int) (*model.Product, error) {
	var p model.Product
	row, err := m.DB.Query("SELECT * FROM products WHERE id=?;", ID)
	if err != nil {
		return nil, err
	}
	row.Next()

	err = row.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Price, &p.ImageURL, &p.Description, &p.Created_at, &p.Updated_at)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (m *ProductSqlMapper) Update(p *model.Product) error {
	q := "UPDATE products SET category_id = ?, title = ?, image_url = ?, price = ?, description = ? WHERE id = ?;"
	_, err := m.DB.Query(q, p.CategoryID, p.Title, p.ImageURL, p.Price, p.Description, p.ID)
	return err
}

func (m *ProductSqlMapper) Delete(ID int) error {
	q := "DELETE FROM products WHERE id = ?;"
	_, err := m.DB.Query(q, ID)
	return err
}

func (m *ProductSqlMapper) List(pn *helper.Pagination) (*[]model.Product, error) {
	var products []model.Product
	q := "SELECT * FROM products ORDER BY id "
	if pn.Order == "asc" {
		q = q + "ASC LIMIT ?,?"
	} else {
		q = q + "DESC LIMIT ?,?"
	}
	row, err := m.DB.Query(q, pn.PerPage*(pn.Page-1), pn.PerPage)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var p model.Product
		err = row.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Price, &p.ImageURL, &p.Description, &p.Created_at, &p.Updated_at)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &products, nil

}
