package utils

import (
	"strings"

	"github.com/shopspring/decimal"
)

func FormatRupiah(d decimal.Decimal) string {
	// Bulatkan ke 2 desimal
	s := d.Round(2).StringFixed(2) // "12345678.90"

	parts := strings.Split(s, ".")
	integer := parts[0]
	fraction := parts[1]

	// Tambahkan pemisah ribuan
	var result []byte
	count := 0

	for i := len(integer) - 1; i >= 0; i-- {
		result = append([]byte{integer[i]}, result...)
		count++
		if count%3 == 0 && i != 0 {
			result = append([]byte{'.'}, result...)
		}
	}

	return "Rp. " + string(result) + "," + fraction
}
