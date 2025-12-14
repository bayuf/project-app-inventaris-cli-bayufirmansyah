package cmd

import (
	"fmt"
	"os"

	categoriescmd "github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/cmd/categories_cmd"
	itemscmd "github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/cmd/items_cmd"
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
		fmt.Println("6. Item Manajemen")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			categoriescmd.GetItemsCategory(h)
		case 2:
			utils.ClearScreen()
			categoriescmd.GetItemsCategoryById(h)
		case 3:
			utils.ClearScreen()
			categoriescmd.AddNewCategory(h)
		case 4:
			utils.ClearScreen()
			categoriescmd.UpdateCategory(h)
		case 5:
			utils.ClearScreen()
			categoriescmd.DeleteCategory(h)
		case 6:
			utils.ClearScreen()
			itemscmd.ItemManagement(h)
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
