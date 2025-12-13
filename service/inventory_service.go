package service

import (
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/repository"
)

type InventoryServiceIface interface {
	GetItemCategory() ([]dto.CategoryResponseDTO, error)
	GetItemByCategoryId(dto.GetItemByCategoryDTO) (dto.CategoryResponseDTO, error)

	AddNewCategory(dto.CreateCategoryDTO) error
	UpdateCategory(dto.UpdateCategoryDTO) error
}

type InventoryService struct {
	repo repository.InventoryIface
}

func NewInventoryService(repo repository.InventoryIface) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

func (s *InventoryService) GetItemCategory() ([]dto.CategoryResponseDTO, error) {
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

	itemCategory, err := s.repo.GetItemByCategoryId(model.Item{ID: item.ID})
	if err != nil {
		return dto.CategoryResponseDTO{}, err
	}

	return dto.CategoryResponseDTO{
		ID:          itemCategory.ID,
		Name:        itemCategory.Name,
		Description: itemCategory.Description,
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
		Name:        *newData.Name,
		Description: *newData.Description,
	}); err != nil {
		return err
	}

	return nil
}
