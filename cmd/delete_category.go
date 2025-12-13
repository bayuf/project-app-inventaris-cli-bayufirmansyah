package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func DeleteCategory(handl handler.InventoryHandler) {
	var categoryId int
OUTTER_LOOP:
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		showAllCategories(handl)

		fmt.Print("Hapus Kategori dengan ID: ")
		fmt.Scan(&categoryId)
		fmt.Println()

		choice := ""
	INNER_LOOP:
		for {
			fmt.Print("Yakin menghapus data ini? (y/n) : ")
			fmt.Scan(&choice)
			fmt.Println()

			choice = strings.ToLower(choice)
			switch choice {
			case "y":
				if err := handl.DeleteCategoryById(categoryId); err != nil {
					utils.ClearScreen()
					fmt.Println(err)

				} else {
					utils.ClearScreen()
					fmt.Println("data with ID", categoryId, "successfully deleted")

				}
			case "n":
				utils.ClearScreen()
				fmt.Println("batal menghapus data dengan ID:", categoryId)
			}

			showAllCategories(handl)

			fmt.Println("1. Hapus Category Lagi")
			fmt.Println("2. Kembali ke Homepage")
			fmt.Println("3. Keluar Aplikasi")
			fmt.Print("Pilih Menu: ")
			choiceInt := 0
			fmt.Scan(&choiceInt)

			switch choiceInt {
			case 1:
				utils.ClearScreen()
				break INNER_LOOP
			case 2:
				utils.ClearScreen()
				break OUTTER_LOOP
			case 3:
				fmt.Println("keluar aplikasi ....")
				os.Exit(0)
			default:
				utils.ClearScreen()
				fmt.Println("Invalid Input")
			}
		}

	}
}
