package itemscmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/utils"
)

func GetTotalInvestmentValue(handl handler.InventoryHandler) {
	fmt.Println("=================== INVENTARIS KANTOR LUMOSHIVE =====================")
	total, err := handl.GetTotalInvestmentValue()
	if err != nil {
		fmt.Println(err)
		return
	}

	showItems(handl)

	fmt.Printf("Total Investment dari semua barang adalah %s\n", utils.FormatRupiah(total.TotalInvestment))
	fmt.Print("Tekan Enter untuk Kembali .....")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	reader.ReadString('\n')
	utils.ClearScreen()
}
