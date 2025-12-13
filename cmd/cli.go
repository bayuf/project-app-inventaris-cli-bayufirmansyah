package cmd

import (
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func Homepage(h handler.InventoryHandler) {
	utils.ClearScreen()
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		fmt.Println("1. Lihat List Kategori Barang")
		fmt.Println("2. Lihat Detail Kategori Barang")
		fmt.Println("3. Tambah Kategori Baru")
		fmt.Println("4. Update Kategori")
		fmt.Println("5. Hapus Kategory")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			GetItemsCategory(h)
		case 2:
			utils.ClearScreen()
			GetItemsCategoryById(h)
		case 3:
			utils.ClearScreen()
			AddNewCategory(h)
		case 4:
			utils.ClearScreen()
			UpdateCategory(h)
		case 5:
			utils.ClearScreen()
			DeleteCategory(h)
		case 6:
			utils.ClearScreen()
			fmt.Println("Keluar dari aplikasi ....")
			os.Exit(0)
		default:
			utils.ClearScreen()
			fmt.Println("invalid input")
		}
	}
}
