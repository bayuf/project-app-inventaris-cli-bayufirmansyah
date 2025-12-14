package categoriescmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func AddNewCategory(handl handler.InventoryHandler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		reader.ReadString('\n')

		fmt.Println("Massukkan Data Kategory Baru")
		fmt.Print("Nama Kategori: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Deskripsi [opsional]: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		if err := handl.AddNewCategory(dto.CreateCategoryDTO{
			Name:        name,
			Description: description,
		}); err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		} else {
			fmt.Printf("Category: %s added successfully\n", name)
		}

		fmt.Println("1. Tambah Category Baru Lagi")
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
