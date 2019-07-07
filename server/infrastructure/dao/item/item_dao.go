package item

import (
	"todo/server/domain/model"
	"todo/server/domain/repository"

	"github.com/jinzhu/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Item {
	return repositoryImpl{
		db: db,
	}
}

func (r repositoryImpl) FirstByInt(id int) *model.Item {
	item := &model.Item{}
	r.db.First(item, id)
	return item
}

func (r repositoryImpl) FindAll() []*model.Item {
	results := []*model.Item{}
	r.db.Find(&results)
	return results
}

func (r repositoryImpl) Create(m *model.Item) *model.Item {
	result := &model.Item{}
	r.db.Create(m).Scan(result)
	return result
}

func (r repositoryImpl) Update(m *model.Item) *model.Item {
	result := &model.Item{}
	r.db.Save(m).Scan(result)
	return result
}
