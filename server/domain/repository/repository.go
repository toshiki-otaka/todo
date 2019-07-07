package repository

import "todo/server/domain/model"

type Item interface {
	FirstByInt(id int) *model.Item
	FindAll() []*model.Item
	Create(*model.Item) *model.Item
	Update(*model.Item) *model.Item
}
