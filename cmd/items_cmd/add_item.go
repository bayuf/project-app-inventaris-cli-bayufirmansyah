package itemscmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func AddItem(handl handler.InventoryHandler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
		showAllCategory(handl)
		categoryId := 0
		fmt.Print("masukkan id category sebelum menambah items ID : ")
		fmt.Scan(&categoryId)
		exists, err := handl.CheckCategory(categoryId)
		if err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		}

		if !exists {
			utils.ClearScreen()
			fmt.Println("category not found")
			continue
		}

		dataItem, _ := handl.GetItemsByCategoryId(categoryId)
		fmt.Println("Data Pada Kategori:", categoryId)
		for _, item := range dataItem {
			fmt.Printf("Categori : %d, Nama Item: %s, Kode SKU: %s\n", item.CategoryId, item.Name, item.SKU)
		}
		fmt.Println("===============================================")

		reader.ReadString('\n')
		fmt.Println("Masukkan Data Item Baru")
		fmt.Print("Nama Item\t\t\t: ")
		name, _ := reader.ReadString('\n')
		validName := utils.NameValidation(name)

		fmt.Print("SKU Item[FOR-123]\t\t: ")
		sku, _ := reader.ReadString('\n')
		validSku := utils.SKUValidation(sku)

		fmt.Print("Harga Item\t\t\t: ")
		price, _ := reader.ReadString('\n')
		validPrice := utils.PriceValidation(price)

		fmt.Print("Tanggal Pembelian[YYYY-MM-DD]\t: ")
		buyDate, _ := reader.ReadString('\n')
		validBuyDate := utils.DateValidation(buyDate)

		fmt.Print("Masa Pakai [Hari]\t\t: ")
		useDays, _ := reader.ReadString('\n')
		validUseDays := utils.DayValidation(useDays)

		fmt.Print("Catatan Tambahan [opsional]\t: ")
		note, _ := reader.ReadString('\n')
		validNote := utils.NoteValidation(note)

		if err := handl.AddNewItem(dto.CreateItemDTO{
			CategoryId: categoryId,
			Name:       validName,
			Price:      validPrice,
			SKU:        validSku,
			BuyDate:    validBuyDate,
			LifeDays:   validUseDays,
			Note:       validNote,
		}); err != nil {
			utils.ClearScreen()
			fmt.Println(err)
			continue
		} else {
			utils.ClearScreen()
			fmt.Printf("Item: %s added successfully\n", validName)
		}

		showItems(handl)
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
