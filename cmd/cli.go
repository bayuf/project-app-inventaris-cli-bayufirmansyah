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
		utils.ClearScreen()
		fmt.Println("============ INVENTARIS KANTOR LUMOSHIVE ==============")
		fmt.Println("1. Lihat List Kategori Barang")
		fmt.Println("2. Lihat Detail Kategori Barang")
		fmt.Println("3. Tambah Kategori Baru")
		fmt.Println("4. Update Kategori")
		fmt.Println("5. Hapus Kategory")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Menu: ")
		choice := 1
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
		case 2:
			utils.ClearScreen()
		case 3:
			utils.ClearScreen()
		case 4:
			utils.ClearScreen()
		case 5:
			utils.ClearScreen()
		case 6:
			utils.ClearScreen()
			fmt.Println("Keluar dari aplikasi ....")
			os.Exit(0)
		}
	}
}
