package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/db"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
	"github.com/shopspring/decimal"
)

type InventoryIface interface {
	// View
	GetItemsCategory() ([]model.ItemCategory, error)
	GetItemCategoryById(model.ItemCategory) (model.ItemCategory, error)
	GetItems() ([]model.Item, error)
	GetItemsByCategoryId(model.Item) ([]model.Item, error)
	GetItemById(model.Item) (model.Item, error)
	CheckCategory(model.ItemCategory) (bool, error)
	GetItemNeedReplacement() ([]model.Item, error)
	GetTotalInvestmentValue() (model.Item, error)
	GetItemDepreciationById(model.Item) (model.Item, error)

	// Add
	AddNewCategory(model.ItemCategory) error
	AddNewItem(model.Item) error

	// Update
	UpdateCategory(model.ItemCategory) error
	UpdateItem(newData model.Item) error

	// Delete
	DeleteCategoryById(model.ItemCategory) error
	DeleteItemById(model.Item) error
}

type Inventory struct {
	DB db.PgxIface
}

func NewInventoryRepository(db db.PgxIface) *Inventory {
	return &Inventory{
		DB: db,
	}
}

// Categories Function
func (i *Inventory) CheckCategory(data model.ItemCategory) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM categories WHERE id = $1 AND deleted_at IS NULL);`
	var exist bool
	if err := i.DB.QueryRow(context.Background(), query, data.ID).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func (i *Inventory) GetItemsCategory() ([]model.ItemCategory, error) {

	query := `SELECT id, name, description FROM categories WHERE deleted_at IS NULL ORDER BY id ASC`
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

func (i *Inventory) GetItemCategoryById(item model.ItemCategory) (model.ItemCategory, error) {
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

func (i *Inventory) GetItemById(data model.Item) (model.Item, error) {
	// query := `SELECT id, name, description FROM categories WHERE id=$1 AND deleted_at IS NULL`
	query := `SELECT id, name, sku, purchase_price, purchase_date, current_status, note
	FROM items WHERE id=$1 AND deleted_at IS NULL`
	row := i.DB.QueryRow(context.Background(), query, data.ID)

	var item model.Item
	if err := row.Scan(&item.ID, &item.Name, &item.SKU, &item.Price, &item.BuyDate, &item.Status, &item.Note); err != nil {
		return model.Item{}, errors.New("error: id not found or not valid")
	}

	return item, nil
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

func (i *Inventory) DeleteCategoryById(item model.ItemCategory) error {
	categoryId := item.ID
	query := `UPDATE categories SET deleted_at=NOW(), updated_at=NOW() WHERE id=$1 AND deleted_at IS NULL RETURNING id`

	var deletedId int
	if err := i.DB.QueryRow(context.Background(), query, categoryId).Scan(&deletedId); err != nil {
		if deletedId == 0 {
			return errors.New("id not found or not valid")
		}

		return err
	}

	return nil
}

// Items Function
func (i *Inventory) GetItems() ([]model.Item, error) {
	query := `SELECT i.id, c.name, i.name, i.purchase_price, i.purchase_date, (CURRENT_DATE - i.purchase_date) AS total_usage_days
FROM items i
LEFT JOIN categories c ON i.category_id = c.id AND c.deleted_at IS NULL
WHERE i.deleted_at IS NULL`

	rows, err := i.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		item := model.Item{}

		rows.Scan(&item.ID, &item.Category, &item.Name, &item.Price, &item.BuyDate, &item.TotalUsage)

		items = append(items, item)
	}

	return items, nil
}

func (i *Inventory) GetItemsByCategoryId(item model.Item) ([]model.Item, error) {
	query := `SELECT category_id, name, sku FROM items WHERE category_id=$1 AND deleted_at IS NULL`
	rows, err := i.DB.Query(context.Background(), query, item.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		item := model.Item{}
		rows.Scan(&item.CategoryId, &item.Name, &item.SKU)

		items = append(items, item)
	}

	return items, nil
}

func (i *Inventory) AddNewItem(newItem model.Item) error {
	query := `INSERT INTO items (category_id, name, sku, purchase_price, purchase_date, useful_life_days, note, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) ON CONFLICT (sku) DO NOTHING;`

	cmdTag, err := i.DB.Exec(context.Background(), query, newItem.CategoryId, newItem.Name, newItem.SKU, newItem.Price, newItem.BuyDate, newItem.LifeDays, newItem.Note)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("item (SKU) already exists")
	}
	return nil
}

func (i *Inventory) DeleteItemById(item model.Item) error {
	id := item.ID
	query := `UPDATE items SET deleted_at=NOW(), updated_at=NOW() WHERE id=$1 AND deleted_at IS NULL RETURNING id`

	var deletedId int
	if err := i.DB.QueryRow(context.Background(), query, id).Scan(&deletedId); err != nil {
		if deletedId == 0 {
			return errors.New("id not found or not valid")
		}

		return err
	}

	return nil
}

func (i *Inventory) UpdateItem(newData model.Item) error {

	query := "UPDATE items SET "
	args := []any{}
	counter := 1

	if newData.Name != "" {
		query += fmt.Sprintf("name=$%d, ", counter)
		args = append(args, newData.Name)
		counter++
	}

	if newData.Note != "" {
		query += fmt.Sprintf("note=$%d, ", counter)
		args = append(args, newData.Note)
		counter++
	}

	if newData.Status != "" {
		query += fmt.Sprintf("current_status=$%d, ", counter)
		args = append(args, newData.Status)
		counter++
	}

	if newData.Name != "" || newData.Note != "" || newData.Status == "" {
		query += "updated_at=now()"
	}
	query = strings.TrimRight(query, ", ")
	query += fmt.Sprintf(" WHERE id=$%d AND deleted_at IS NULL RETURNING id", counter)
	args = append(args, newData.ID)

	var UpdatedId int
	if err := i.DB.QueryRow(context.Background(), query, args...).Scan(&UpdatedId); err != nil {
		if UpdatedId == 0 {
			return errors.New("update failed: duplicate name or item not found")
		}

		return err
	}

	return nil
}

func (i *Inventory) GetItemNeedReplacement() ([]model.Item, error) {
	query := `SELECT
    i.id,
    i.name,
    i.sku,
    i.purchase_date,
    (CURRENT_DATE - i.purchase_date) AS usage_days
FROM items i
WHERE i.deleted_at IS NULL
  AND (CURRENT_DATE - i.purchase_date) > 100
ORDER BY usage_days DESC;
`
	rows, err := i.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		item := model.Item{}
		rows.Scan(&item.ID, &item.Name, &item.SKU, &item.BuyDate, &item.LifeDays)

		items = append(items, item)
	}

	return items, nil
}

func (i *Inventory) GetTotalInvestmentValue() (model.Item, error) {
	query := `SELECT
    SUM(
        i.purchase_price *
        POWER(
            1 - i.depreciation_rate,
            (CURRENT_DATE - i.purchase_date) / 365.0
        )
    ) AS total_current_value
FROM items i
WHERE i.deleted_at IS NULL;`

	row := i.DB.QueryRow(context.Background(), query)

	var totalInvest decimal.Decimal
	err := row.Scan(&totalInvest)
	if err != nil {
		return model.Item{}, err
	}

	return model.Item{TotalInvestment: totalInvest}, nil
}

func (i *Inventory) GetItemDepreciationById(item model.Item) (model.Item, error) {
	query := `SELECT
    i.id,
    i.name,

    i.purchase_price AS original_value,

    (CURRENT_DATE - i.purchase_date) AS usage_days,
    (CURRENT_DATE - i.purchase_date) / 365.0 AS age_years,

    i.purchase_price *
    POWER(
        1 - i.depreciation_rate,
        (CURRENT_DATE - i.purchase_date) / 365.0
    ) AS current_value,

    i.purchase_price -
    (
        i.purchase_price *
        POWER(
            1 - i.depreciation_rate,
            (CURRENT_DATE - i.purchase_date) / 365.0
        )
    ) AS accumulated_depreciation

FROM items i
WHERE i.id = $1
  AND i.deleted_at IS NULL;`

	row := i.DB.QueryRow(context.Background(), query, item.ID)
	itemResponse := model.Item{}
	if err := row.Scan(&itemResponse.ID, &itemResponse.Name, &itemResponse.Price, &itemResponse.TotalUsage, &itemResponse.AgeItem, &itemResponse.CurrentValue, &itemResponse.DepreciationValue); err != nil {
		return model.Item{}, nil
	}

	return itemResponse, nil

}
