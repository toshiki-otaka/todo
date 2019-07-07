package usecase

import (
	"time"
	"todo/server/domain/model"
	"todo/server/domain/repository"
)

type ItemUseCase struct {
	repo repository.Item
}

func NewItemUseCase(repo repository.Item) ItemUseCase {
	return ItemUseCase{
		repo: repo,
	}
}

func (uc ItemUseCase) GetItemList() []*model.Item {
	return uc.repo.FindAll()
}

func (uc ItemUseCase) Create(title, desc string) *model.Item {
	item := &model.Item{
		Title:      title,
		Desc:       desc,
		InsertedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}

	return uc.repo.Create(item)
}

func (uc ItemUseCase) Close(itemId int) *model.Item {
	item := uc.repo.FirstByInt(itemId)
	item.IsDone = true

	return uc.repo.Update(item)
}

func (uc ItemUseCase) Open(itemId int) *model.Item {
	item := uc.repo.FirstByInt(itemId)
	item.IsDone = false

	return uc.repo.Update(item)
}
