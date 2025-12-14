package utils

import (
	"fmt"
	"os"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/dto"
	"github.com/olekukonko/tablewriter"
)

func PrintTableCategory(itemsCategory []dto.CategoryResponseDTO) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"No", "ID", "Kategori", "Deskripsi Kategori"})

	for i, t := range itemsCategory {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%d", t.ID),
			t.Name,
			t.Description,
		}

		table.Append(row)
	}
	table.Render()
}

func PrintTableCategoryById(itemsCategory []dto.CategoryResponseDTO) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"No", "ID", "Kategori", "Deskripsi Kategori", "Quantity"})

	for i, t := range itemsCategory {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%d", t.ID),
			t.Name,
			t.Description,
			fmt.Sprintf("%d", t.Quantity),
		}

		table.Append(row)
	}
	table.Render()
}

func PrintTableItems(items []dto.ItemResponseDTO) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"No", "ID", "Kategori", "Nama", "Harga", "Tanggal Beli", "Umur Penggunaan"})

	for i, t := range items {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%d", t.ID),
			t.Category,
			t.Name,
			fmt.Sprintf("Rp. %v", t.Price),
			GetDateFormat(t.BuyDate),
			fmt.Sprintf("%d Hari", t.TotalUsage),
		}

		table.Append(row)
	}
	table.Render()
}
