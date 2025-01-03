package services

import (
	"gin-flemarket/dto"
	"gin-flemarket/models"
	"gin-flemarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, UpdateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemService struct {
	repositories repositories.IItemRepository
}

func NewItemService(repositories repositories.IItemRepository) IItemService {
	return &ItemService{repositories: repositories}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repositories.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repositories.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}

	return s.repositories.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}

	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repositories.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint) error {
	return s.repositories.Delete(itemId)
}
