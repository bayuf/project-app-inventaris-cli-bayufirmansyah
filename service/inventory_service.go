package service

import (
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/repository"
)

type InventoryServiceIface interface {
	GetItemsCategory() ([]dto.CategoryResponseDTO, error)
	GetItemByCategoryId(dto.GetItemByCategoryDTO) (dto.CategoryResponseDTO, error)

	AddNewCategory(dto.CreateCategoryDTO) error
	UpdateCategory(dto.UpdateCategoryDTO) error
	DeleteCategoryById(dto.UpdateCategoryDTO) error

	// Item
	GetItems() ([]dto.ItemResponseDTO, error)
}

type InventoryService struct {
	repo repository.InventoryIface
}

func NewInventoryService(repo repository.InventoryIface) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

// Categories Function
func (s *InventoryService) GetItemsCategory() ([]dto.CategoryResponseDTO, error) {
	items, err := s.repo.GetItemsCategory()
	if err != nil {
		return nil, err
	}

	// convert from model to dto
	var itemsCategory []dto.CategoryResponseDTO
	for _, v := range items {
		itemCategory := dto.CategoryResponseDTO{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		}

		itemsCategory = append(itemsCategory, itemCategory)
	}

	return itemsCategory, nil
}

func (s *InventoryService) GetItemByCategoryId(item dto.GetItemByCategoryDTO) (dto.CategoryResponseDTO, error) {

	itemCategory, err := s.repo.GetItemByCategoryId(model.ItemCategory{ID: item.ID})
	if err != nil {
		return dto.CategoryResponseDTO{}, err
	}

	return dto.CategoryResponseDTO{
		ID:          itemCategory.ID,
		Name:        itemCategory.Name,
		Description: itemCategory.Description,
		Quantity:    itemCategory.Quantity,
	}, nil
}

func (s *InventoryService) AddNewCategory(newCategory dto.CreateCategoryDTO) error {
	if err := s.repo.AddNewCategory(model.ItemCategory{Name: newCategory.Name, Description: newCategory.Description}); err != nil {
		return err
	}
	return nil
}

func (s *InventoryService) UpdateCategory(newData dto.UpdateCategoryDTO) error {

	if err := s.repo.UpdateCategory(model.ItemCategory{
		ID:          newData.ID,
		Name:        newData.Name,
		Description: newData.Description,
	}); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) DeleteCategoryById(item dto.UpdateCategoryDTO) error {
	if err := s.repo.DeleteCategoryById(model.ItemCategory{ID: item.ID}); err != nil {
		return err
	}

	return nil
}

// Item Function
func (s *InventoryService) GetItems() ([]dto.ItemResponseDTO, error) {
	items, err := s.repo.GetItems()
	if err != nil {
		return nil, err
	}

	itemsResponse := []dto.ItemResponseDTO{}
	for _, v := range items {
		item := dto.ItemResponseDTO{
			ID:        v.ID,
			Name:      v.Name,
			Price:     v.Price,
			BuyDate:   v.BuyDate,
			TotalUsed: v.TotalUsed,
		}
		itemsResponse = append(itemsResponse, item)
	}
	return itemsResponse, nil
}
