package repository

import (
	"context"
	"errors"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/db"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
)

type InventoryIface interface {
	GetItemsCategory() ([]model.ItemCategory, error)
	GetItemByCategoryId(model.Item) (model.ItemCategory, error)
}

type Inventory struct {
	DB db.PgxIface
}

func NewInventoryRepository(db db.PgxIface) *Inventory {
	return &Inventory{
		DB: db,
	}
}

func (i *Inventory) GetItemsCategory() ([]model.ItemCategory, error) {

	query := `SELECT id, name, description FROM categories WHERE deleted_at IS NULL`
	rows, err := i.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ItemsCategory []model.ItemCategory
	for rows.Next() {
		var ItemCategory model.ItemCategory
		if err := rows.Scan(&ItemCategory.ID, &ItemCategory.Name, &ItemCategory.Description); err != nil {
			return nil, err
		}
		ItemsCategory = append(ItemsCategory, ItemCategory)
	}

	return ItemsCategory, nil
}

func (i *Inventory) GetItemByCategoryId(item model.Item) (model.ItemCategory, error) {
	itemCategoryId := item.ID

	query := `SELECT id, name, description FROM categories WHERE id=$1 AND deleted_at IS NULL`
	row := i.DB.QueryRow(context.Background(), query, itemCategoryId)

	var ItemCategory model.ItemCategory
	if err := row.Scan(&ItemCategory.ID, &ItemCategory.Name, &ItemCategory.Description); err != nil {
		return model.ItemCategory{}, errors.New("error: id not found or not valid")
	}

	return ItemCategory, nil
}
