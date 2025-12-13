package cmd

import (
	"fmt"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
)

func Homepage(h handler.InventoryHandler) {

	// item, err := h.GetItemByCategoryId(12)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("ID : %d Name: %s Description: % s", item.ID, item.Name, item.Description)
	// }

	// items, err := h.GetItemCategory()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("ID---|Category---------|Description-------------|")
	// 	for _, item := range items {
	// 		fmt.Printf("%d   |%s            |%s                      |\n", item.ID, item.Name, item.Description)
	// 	}
	// }

	if err := h.AddNewCategory(dto.CreateCategoryDTO{Name: "monitor", Description: "Test Description"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add succeed")
	}
}
