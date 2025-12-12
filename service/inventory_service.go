package service

import "github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/repository"

type InventoryServiceIface interface {
}

type InventoryService struct {
	repo repository.InventoryIface
}

func NewInventoryService(repo repository.InventoryIface) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}
