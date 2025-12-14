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
func (h *InventoryHandler) CheckCategory(id int) (bool, error) {
	exist, err := h.service.CheckCategory(id)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *InventoryHandler) GetItemsCategory() ([]dto.CategoryResponseDTO, error) {
	return h.service.GetItemsCategory()
}

func (h *InventoryHandler) GetItemCategoryById(id int) (dto.CategoryResponseDTO, error) {
	item, err := h.service.GetItemCategoryById(dto.GetItemByCategoryDTO{ID: id})
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

func (h *InventoryHandler) GetItemById(id int) (dto.ItemResponseDTO, error) {
	items, err := h.service.GetItemById(dto.GetItemDTO{ItemID: id})
	if err != nil {
		return dto.ItemResponseDTO{}, err
	}

	return items, nil
}
func (h *InventoryHandler) GetItemsByCategoryId(id int) ([]dto.ItemResponseDTO, error) {
	items, err := h.service.GetItemsByCategoryId(dto.GetItemDTO{ItemID: id})
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (h *InventoryHandler) AddNewItem(data dto.CreateItemDTO) error {
	if err := h.service.AddNewItem(dto.CreateItemDTO{
		CategoryId: data.CategoryId,
		Name:       data.Name,
		SKU:        data.SKU,
		Price:      data.Price,
		BuyDate:    data.BuyDate,
		LifeDays:   data.LifeDays,
		Note:       data.Note,
	}); err != nil {
		return err
	}

	return nil
}

func (h *InventoryHandler) DeleteItemById(id int) error {
	if err := h.service.DeleteItemById(dto.UpdateItemDTO{ID: id}); err != nil {
		return err
	}
	return nil
}

func (h *InventoryHandler) UpdateItem(newData dto.UpdateItemDTO) error {
	if err := h.service.UpdateItem(newData); err != nil {
		return err
	}

	return nil
}
