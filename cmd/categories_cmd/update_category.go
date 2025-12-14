package categoriescmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func UpdateCategory(handl handler.InventoryHandler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")

		showAllCategories(handl)

		fmt.Print("Update Data Kategori ID: ")
		var categoryId int
		fmt.Scan(&categoryId)

		utils.ClearScreen()
		item, err := handl.GetItemByCategoryId(categoryId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		utils.PrintTableCategory([]dto.CategoryResponseDTO{{ID: item.ID, Name: item.Name, Description: item.Description}})
		reader.ReadString('\n')

		fmt.Println("Masukkan Data Baru")
		fmt.Print("Nama Baru: ")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)

		fmt.Print("Deskripsi Baru [opsional]: ")
		newDescription, _ := reader.ReadString('\n')
		newDescription = strings.TrimSpace(newDescription)

		if err := handl.UpdateCategory(dto.UpdateCategoryDTO{
			ID:          categoryId,
			Name:        newName,
			Description: newDescription,
		}); err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		} else {
			utils.ClearScreen()
			fmt.Println("Category: update successfully")
		}

		showAllCategories(handl)

		fmt.Println("1. Update Category Lagi")
		fmt.Println("2. Kembali ke Homepage")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			continue
		case 2:
			utils.ClearScreen()
			return
		case 3:
			fmt.Println("keluar aplikasi ....")
			os.Exit(0)
		default:
			utils.ClearScreen()
			fmt.Println("Invalid Input")
		}

	}
}
