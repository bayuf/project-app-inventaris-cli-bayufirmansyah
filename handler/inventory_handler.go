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

func (h *InventoryHandler) GetItemCategory() ([]dto.CategoryResponseDTO, error) {
	return h.service.GetItemCategory()
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
	}, nil
}

func (h *InventoryHandler) AddNewCategory(newCategoryData dto.CreateCategoryDTO) error {
	newCategory := newCategoryData
	if err := h.service.AddNewCategory(newCategory); err != nil {
		return err
	}

	return nil
}
