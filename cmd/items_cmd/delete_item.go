package itemscmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func DeleteItem(handl handler.InventoryHandler) {
	var itemId int
OUTTER_LOOP:
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		showItems(handl)

		fmt.Print("Hapus Barang dengan ID: ")
		fmt.Scan(&itemId)
		fmt.Println()

		choice := ""
	INNER_LOOP:
		for {
			fmt.Print("Yakin menghapus Barang ini? (y/n) : ")
			fmt.Scan(&choice)
			fmt.Println()

			choice = strings.ToLower(choice)
			switch choice {
			case "y":
				if err := handl.DeleteItemById(itemId); err != nil {
					utils.ClearScreen()
					fmt.Println(err)

				} else {
					utils.ClearScreen()
					fmt.Println("data with ID", itemId, "successfully deleted")

				}
			case "n":
				utils.ClearScreen()
				fmt.Println("batal menghapus data dengan ID:", itemId)
			}

			showItems(handl)

			fmt.Println("1. Hapus Barang Lagi")
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
