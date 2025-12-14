package itemscmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetDepreciationById(handl handler.InventoryHandler) {
	fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
	showItems(handl)
	fmt.Print("Masukkan ID untuk mengetahui nilai Depreciation: ")
	id := 0
	fmt.Scan(&id)

	i, err := handl.GetItemDepreciationById(dto.GetItemDTO{ItemID: id})
	if err != nil {
		utils.ClearScreen()
		fmt.Println(err)
		return
	}

	utils.ClearScreen()
	utils.PrintTableItemDepreciation([]dto.ItemResponseDTO{
		{
			ID:                i.ID,
			Name:              i.Name,
			Price:             i.Price,
			TotalUsage:        i.TotalUsage,
			AgeItem:           i.AgeItem,
			CurrentValue:      i.CurrentValue,
			DepreciationValue: i.DepreciationValue,
		},
	})
	fmt.Print("Tekan Enter untuk Kembali ....")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	reader.ReadString('\n')
	utils.ClearScreen()
}
