package handler

import (
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/service"
)

type InventoryHandler struct {
	service service.InventoryServiceIface
}

func NewInventoryHandler(service service.InventoryServiceIface) InventoryHandler {
	return InventoryHandler{service: service}
}

// Category Function
func (h *InventoryHandler) GetItemsCategory() ([]dto.CategoryResponseDTO, error) {
	return h.service.GetItemsCategory()
}

func (h *InventoryHandler) GetItemByCategoryId(id int) (dto.CategoryResponseDTO, error) {
	item, err := h.service.GetItemByCategoryId(dto.GetItemByCategoryDTO{ID: id})
	if err != nil {
		return dto.CategoryResponseDTO{}, err
	}

	return dto.CategoryResponseDTO{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
	}, nil
}

func (h *InventoryHandler) AddNewCategory(newCategoryData dto.CreateCategoryDTO) error {
	newCategory := newCategoryData
	if err := h.service.AddNewCategory(newCategory); err != nil {
		return err
	}

	return nil
}

func (h *InventoryHandler) UpdateCategory(newData dto.UpdateCategoryDTO) error {
	if err := h.service.UpdateCategory(newData); err != nil {
		return err
	}

	return nil
}

func (h *InventoryHandler) DeleteCategoryById(id int) error {
	if err := h.service.DeleteCategoryById(dto.UpdateCategoryDTO{ID: id}); err != nil {
		return err
	}
	return nil
}

// Item Function
func (h *InventoryHandler) GetItems() ([]dto.ItemResponseDTO, error) {
	items, err := h.service.GetItems()
	if err != nil {
		return nil, err
	}

	return items, nil
}
