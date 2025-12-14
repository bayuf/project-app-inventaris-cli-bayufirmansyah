package itemscmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetItems(handl handler.InventoryHandler) {
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		showItems(handl)
		fmt.Println("1. Lihat Detail Barang")
		fmt.Println("2. Kembali ke Item Manajemen")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
			showItems(handl)
			fmt.Print("Masukkan Id Barang :")
			id := 0
			fmt.Scan(&id)
			item, err := handl.GetItemById(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			utils.ClearScreen()
			utils.PrintTableItem([]dto.ItemResponseDTO{{
				ID:      item.ID,
				Name:    item.Name,
				SKU:     item.SKU,
				Price:   item.Price,
				BuyDate: item.BuyDate,
				Status:  item.Status,
				Note:    item.Note,
			}})
			fmt.Print("tekan enter untuk kembali .....")
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')
			reader.ReadString('\n')
			utils.ClearScreen()
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
