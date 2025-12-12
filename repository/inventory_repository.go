package repository

import (
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/db"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/model"
)

type InventoryIface interface {
	GetItemById(model.Item) (model.Item, error)
}

type Inventory struct {
	DB db.PgxIface
}

func NewInventoryRepository(db db.PgxIface) *Inventory {
	return &Inventory{
		DB: db,
	}
}

func (i *Inventory) GetItemById(item model.Item) (model.Item, error) {
	return model.Item{}, nil
}
