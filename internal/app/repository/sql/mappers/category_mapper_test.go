package mappers

import (
	"bestprice_test/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertCategory(t *testing.T) {
	c := model.Category{}
	cm := CategorySqlMapper{}

	err := cm.Insert(&c)
	assert.Nil(t, err)
}

func TestReadCategory(t *testing.T) {
	ID := 1
	cm := CategorySqlMapper{}

	actual, err := cm.Read(ID)
	assert.Nil(t, err)
	assert.Equal(t, actual.ID, ID)
}

func TestUpdateCategory(t *testing.T) {
	c := model.Category{}
	cm := CategorySqlMapper{}

	err := cm.Update(&c)
	assert.Nil(t, err)
}

func TestDeleteCategory(t *testing.T) {
	ID := 1
	cm := CategorySqlMapper{}

	err := cm.Delete(ID)
	assert.Nil(t, err)
}
