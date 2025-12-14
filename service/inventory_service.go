package service

import (
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/repository"
)

type InventoryServiceIface interface {
	CheckCategory(id int) (bool, error)
	GetItemsCategory() ([]dto.CategoryResponseDTO, error)
	GetItemCategoryById(dto.GetItemByCategoryDTO) (dto.CategoryResponseDTO, error)

	AddNewCategory(dto.CreateCategoryDTO) error
	UpdateCategory(dto.UpdateCategoryDTO) error
	DeleteCategoryById(dto.UpdateCategoryDTO) error

	// Item
	GetItems() ([]dto.ItemResponseDTO, error)
	DeleteItemById(dto.UpdateItemDTO) error
	AddNewItem(dto.CreateItemDTO) error
	GetItemsByCategoryId(dto.GetItemDTO) ([]dto.ItemResponseDTO, error)
	GetItemById(dto.GetItemDTO) (dto.ItemResponseDTO, error)
	UpdateItem(newData dto.UpdateItemDTO) error
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
func (s *InventoryService) CheckCategory(id int) (bool, error) {

	exist, err := s.repo.CheckCategory(model.ItemCategory{ID: id})
	if err != nil {
		return exist, err
	}

	return exist, nil
}

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

func (s *InventoryService) GetItemCategoryById(item dto.GetItemByCategoryDTO) (dto.CategoryResponseDTO, error) {

	itemCategory, err := s.repo.GetItemCategoryById(model.ItemCategory{ID: item.ID})
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
			ID:         v.ID,
			Category:   v.Category,
			Name:       v.Name,
			Price:      v.Price,
			BuyDate:    v.BuyDate,
			TotalUsage: v.TotalUsage,
		}
		itemsResponse = append(itemsResponse, item)
	}
	return itemsResponse, nil
}

func (s *InventoryService) GetItemById(id dto.GetItemDTO) (dto.ItemResponseDTO, error) {
	item, err := s.repo.GetItemById(model.Item{ID: id.ItemID})
	if err != nil {
		return dto.ItemResponseDTO{}, err
	}

	return dto.ItemResponseDTO{
		ID:      item.ID,
		Name:    item.Name,
		SKU:     item.SKU,
		Price:   item.Price,
		BuyDate: item.BuyDate,
		Status:  item.Status,
		Note:    item.Note,
	}, nil

}
func (s *InventoryService) GetItemsByCategoryId(id dto.GetItemDTO) ([]dto.ItemResponseDTO, error) {
	items, err := s.repo.GetItemsByCategoryId(model.Item{CategoryId: id.ItemID})
	if err != nil {
		return nil, err
	}

	itemResponse := []dto.ItemResponseDTO{}
	for _, v := range items {
		item := dto.ItemResponseDTO{
			CategoryId: v.CategoryId,
			Name:       v.Name,
			SKU:        v.SKU,
		}
		itemResponse = append(itemResponse, item)
	}

	return itemResponse, nil

}

func (s *InventoryService) AddNewItem(newItem dto.CreateItemDTO) error {

	s.repo.AddNewItem(model.Item{
		CategoryId: newItem.CategoryId,
		Name:       newItem.Name,
		SKU:        newItem.SKU,
		Price:      newItem.Price,
		BuyDate:    newItem.BuyDate,
		LifeDays:   newItem.LifeDays,
		Note:       newItem.Note,
	})
	return nil
}

func (s *InventoryService) DeleteItemById(item dto.UpdateItemDTO) error {
	if err := s.repo.DeleteItemById(model.Item{ID: item.ID}); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) UpdateItem(newData dto.UpdateItemDTO) error {

	if err := s.repo.UpdateItem(model.Item{
		ID:     newData.ID,
		Name:   newData.Name,
		Note:   newData.Note,
		Status: newData.Status,
	}); err != nil {
		return err
	}

	return nil
}
