package mappers

import "bestprice_test/internal/app/model"

type SqlMapper interface {
	Insert(c *model.Category) error
	Read(ID int) (*model.Category, error)
	Update(c *model.Category) error
	Delete(ID int) error
	List()
}
