package mappers
import (
	"bestprice_test/internal/app/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CategorySqlMapper struct {
	DB *sql.DB
}

func (m *CategorySqlMapper) Insert(c *model.Category) error {
	return nil
}

func (m *CategorySqlMapper) Read(ID int) (*model.Category, error) {
	// var c model.Category
	// test := m.DB.Query("SELECT * FROM category WHERE id=?", ID).Scan(&c.ID, &c.Name)
	// test2 := m.DB.Prepare()
	return &model.Category{ID: ID}, nil
}

func (m *CategorySqlMapper) Update(c *model.Category) (*model.Category, error) {
	return c, nil
}

func (m *CategorySqlMapper) Delete(ID int) error {
	return nil
}

func (m *CategorySqlMapper) List() {

}
