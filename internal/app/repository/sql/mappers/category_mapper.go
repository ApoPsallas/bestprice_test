package mappers

import (
	"bestprice_test/internal/app/helper"
	"bestprice_test/internal/app/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CategorySqlMapper struct {
	DB *sql.DB
}

func (m *CategorySqlMapper) Insert(c *model.Category) error {
	q := "INSERT INTO categories (title,position,image_url) VALUES (?,?,?);"
	_, err := m.DB.Query(q, c.Title, c.Position, c.ImageURL)
	return err
}

func (m *CategorySqlMapper) Read(ID int) (*model.Category, error) {
	var c model.Category
	row, err := m.DB.Query("SELECT * FROM categories WHERE id=?;", ID)
	if err != nil {
		return nil, err
	}
	row.Next()

	err = row.Scan(&c.ID, &c.Title, &c.Position, &c.ImageURL, &c.Created_at, &c.Updated_at)
	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (m *CategorySqlMapper) Update(c *model.Category) error {
	q := "UPDATE categories SET title = ?, position = ?, image_url = ?  WHERE id = ?;"
	_, err := m.DB.Query(q, c.Title, c.Position, c.ImageURL, c.ID)
	return err
}

func (m *CategorySqlMapper) Delete(ID int) error {
	q := "DELETE FROM categories WHERE id = ?;"
	_, err := m.DB.Query(q, ID)
	return err
}

func (m *CategorySqlMapper) List(pn *helper.Pagination) (*[]model.Category, error) {
	var categories []model.Category
	q := "SELECT * FROM categories ORDER BY id "
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
		var c model.Category
		err = row.Scan(&c.ID, &c.Title, &c.Position, &c.ImageURL, &c.Created_at, &c.Updated_at)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return &categories, nil

}
