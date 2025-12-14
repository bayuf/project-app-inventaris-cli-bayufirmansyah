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
		fmt.Println("2. Tambah Item Baru")
		fmt.Println("3. Edit Item")
		fmt.Println("4. Hapus Item")
		fmt.Println("5. Total Investment")
		fmt.Println("6. Harga Depreciation")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			GetItems(handl)
		case 2:
			utils.ClearScreen()
			AddItem(handl)
		case 3:
			utils.ClearScreen()
			UpdateItem(handl)
		case 4:
			utils.ClearScreen()
			DeleteItem(handl)
		case 5:
			utils.ClearScreen()
			GetTotalInvestmentValue(handl)
		case 6:
			utils.ClearScreen()
			GetDepreciationById(handl)
		case 7:
			utils.ClearScreen()
			fmt.Println("Keluar dari aplikasi ....")
			os.Exit(0)
		default:
			utils.ClearScreen()
			fmt.Println("invalid input")
		}
	}
}
