package itemscmd

import (
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func ItemManagement(handl handler.InventoryHandler) {
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		fmt.Println("1. Lihat List Kategori Barang")
		fmt.Println("2. Item Manajemen")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			GetItems(handl)
		case 3:
			utils.ClearScreen()
			fmt.Println("Keluar dari aplikasi ....")
			os.Exit(0)
		default:
			utils.ClearScreen()
			fmt.Println("invalid input")
		}
	}
}
