package itemscmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetItems(handl handler.InventoryHandler) {
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		showItems(handl)
		fmt.Println("1. Kembali ke Item Manajemen")
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

func showItems(handl handler.InventoryHandler) {
	items, err := handl.GetItems()
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.PrintTableItems(items)
}

func showAllCategory(handl handler.InventoryHandler) {
	items, err := handl.GetItemsCategory()
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintTableCategory(items)
}
