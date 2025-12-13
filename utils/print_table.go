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
