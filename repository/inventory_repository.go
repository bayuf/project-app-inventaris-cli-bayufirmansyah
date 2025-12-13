package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/db"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
)

type InventoryIface interface {
	// View
	GetItemsCategory() ([]model.ItemCategory, error)
	GetItemByCategoryId(model.ItemCategory) (model.ItemCategory, error)

	// Add
	AddNewCategory(model.ItemCategory) error

	// Update
	UpdateCategory(model.ItemCategory) error

	// Delete
	DeleteCategoryById(model.ItemCategory) error
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

func (i *Inventory) GetItemByCategoryId(item model.ItemCategory) (model.ItemCategory, error) {
	itemCategoryId := item.ID

	// query := `SELECT id, name, description FROM categories WHERE id=$1 AND deleted_at IS NULL`
	query := `SELECT 
    c.id,
    c.name,
    c.description,
    COUNT(i.id) AS total_items
FROM categories c
LEFT JOIN items i ON i.category_id = c.id AND i.deleted_at IS NULL
WHERE c.id = $1
  AND c.deleted_at IS NULL
GROUP BY c.id, c.name, c.description;`
	row := i.DB.QueryRow(context.Background(), query, itemCategoryId)

	var ItemCategory model.ItemCategory
	if err := row.Scan(&ItemCategory.ID, &ItemCategory.Name, &ItemCategory.Description, &ItemCategory.Quantity); err != nil {
		return model.ItemCategory{}, errors.New("error: id not found or not valid")
	}

	return ItemCategory, nil
}

func (i *Inventory) AddNewCategory(newCategory model.ItemCategory) error {
	query := `INSERT INTO categories (name, description) VALUES ($1, $2) ON CONFLICT DO NOTHING;`

	cmdTag, err := i.DB.Exec(context.Background(), query, newCategory.Name, newCategory.Description)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("category already exist")
	}

	return nil
}

func (i *Inventory) UpdateCategory(newData model.ItemCategory) error {

	query := "UPDATE categories SET "
	args := []any{}
	counter := 1

	if newData.Name != "" {
		query += fmt.Sprintf("name=$%d, ", counter)
		args = append(args, newData.Name)
		counter++
	}

	if newData.Description != "" {
		query += fmt.Sprintf("description=$%d, ", counter)
		args = append(args, newData.Description)
		counter++
	}

	if newData.Name != "" || newData.Description != "" {
		query += "updated_at=now()"
	}
	query = strings.TrimRight(query, ", ")
	query += fmt.Sprintf(" WHERE id=$%d AND deleted_at IS NULL RETURNING id", counter)
	args = append(args, newData.ID)

	var UpdatedId int
	if err := i.DB.QueryRow(context.Background(), query, args...).Scan(&UpdatedId); err != nil {
		if UpdatedId == 0 {
			return errors.New("update failed: duplicate name or category not found")
		}

		return err
	}

	return nil
}

func (i *Inventory) DeleteCategoryById(id model.ItemCategory) error {
	return nil
}
