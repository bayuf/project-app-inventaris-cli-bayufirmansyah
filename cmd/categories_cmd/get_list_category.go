package categoriescmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetItemsCategory(handl handler.InventoryHandler) {
	items, err := handl.GetItemsCategory()
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		utils.PrintTableCategory(items)
		fmt.Println("1. Kembali ke Homepage")
		fmt.Println("2. Keluar Aplikasi")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			return
		case 2:
			fmt.Println("keluar aplikasi ....")
			os.Exit(0)
		default:
			utils.ClearScreen()
			fmt.Println("Invalid Input")
		}

	}

}

func showAllCategories(handl handler.InventoryHandler) {
	items, err := handl.GetItemsCategory()
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintTableCategory(items)
}
