package itemscmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func UpdateItem(handl handler.InventoryHandler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")

		showItems(handl)

		fmt.Print("Update Data ID: ")
		var itemId int
		fmt.Scan(&itemId)

		utils.ClearScreen()
		item, err := handl.GetItemById(itemId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		utils.PrintTableItem([]dto.ItemResponseDTO{
			{
				ID:      item.ID,
				Name:    item.Name,
				SKU:     item.SKU,
				Price:   item.Price,
				BuyDate: item.BuyDate,
				Status:  item.Status,
				Note:    item.Note,
			},
		})
		reader.ReadString('\n')

		fmt.Println("Masukkan Data Baru")
		fmt.Print("Nama\t\t\t: ")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)

		fmt.Print("Catatan\t\t\t: ")
		newNote, _ := reader.ReadString('\n')
		newNote = strings.TrimSpace(newNote)

		fmt.Print("Status[active, retired, maintenance]\t: ")
		newStatus, _ := reader.ReadString('\n')
		newStatus = strings.TrimSpace(newStatus)

		if err := handl.UpdateItem(dto.UpdateItemDTO{
			ID:     itemId,
			Name:   newName,
			Note:   newNote,
			Status: newStatus,
		}); err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		} else {
			utils.ClearScreen()
			fmt.Println("update successfully")
		}

		showItems(handl)

		fmt.Println("1. Update Barang Lagi")
		fmt.Println("2. Kembali ke Homepage")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Pilih Menu: ")
		choice := 0
		fmt.Scan(&choice)

		switch choice {
		case 1:
			utils.ClearScreen()
			continue
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
