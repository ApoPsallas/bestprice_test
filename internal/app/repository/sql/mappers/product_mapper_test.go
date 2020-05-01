package mappers

import (
	"bestprice_test/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertProduct(t *testing.T) {
	p := model.Product{}
	pm := ProductSqlMapper{}

	err := pm.Insert(&p)
	assert.Nil(t, err)
}

func TestReadProduct(t *testing.T) {
	ID := 1
	pm := ProductSqlMapper{}

	actual, err := pm.Read(ID)
	assert.Nil(t, err)
	assert.Equal(t, actual.ID, ID)
}

func TestUpdateProduct(t *testing.T) {
	p := model.Product{}
	pm := ProductSqlMapper{}

	actual, err := pm.Update(&p)
	assert.Nil(t, err)
	assert.Equal(t, actual, &p)
}

func TestDeleteProduct(t *testing.T) {
	ID := 1
	pm := ProductSqlMapper{}

	err := pm.Delete(ID)
	assert.Nil(t, err)
}
