package categoriescmd

import (
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetItemsCategoryById(handl handler.InventoryHandler) {
	categoriId := 0
	for {

		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		fmt.Print("Masukkan Id dari categori barang: ")
		fmt.Scan(&categoriId)
		item, err := handl.GetItemCategoryById(categoriId)
		if err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		}

		utils.PrintTableCategoryById([]dto.CategoryResponseDTO{
			{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				Quantity:    item.Quantity,
			},
		})

		fmt.Println("1. Lihat Category yang lain")
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
