package handler

import "github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/service"

type InventoryHandler struct {
	service service.InventoryServiceIface
}

func NewInventoryHandler(service service.InventoryServiceIface) InventoryHandler {
	return InventoryHandler{service: service}
}
