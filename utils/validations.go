package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func NameValidation(name string) string {
	name = strings.TrimSpace(name)
	name = regexp.MustCompile(`\s+`).ReplaceAllString(name, " ")

	if name == "" {
		log.Fatal("name is empty")
	}

	if len(name) < 3 || len(name) > 50 {
		log.Fatal("name must be 3 - 50 characters")
	}

	pattern := `^[A-Za-z0-9 ]+$`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(name) {
		log.Fatal("invalid name")
	}

	return name

}

func PriceValidation(price string) decimal.Decimal {
	if price == "" {
		log.Fatal("price is empty")
	}
	price = strings.TrimSpace(price)
	validPrice, err := decimal.NewFromString(price)
	if err != nil {
		log.Fatal("price is not number")
	}

	return validPrice
}

func SKUValidation(sku string) string {
	sku = strings.TrimSpace(sku)
	if sku == "" {
		log.Fatal("sku is empty")
	}

	pattern := `^[A-Z]{1,5}-[0-9]{1,10}$`

	sku = strings.ToUpper(sku)
	re := regexp.MustCompile(pattern)

	if !re.MatchString(sku) {
		return "invalid sku"
	}
	return sku
}

func DateValidation(date string) time.Time {
	if date == "" {
		log.Fatal("date is empty")
	}
	date = strings.TrimSpace(date)
	validDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}

	return validDate
}

func DayValidation(days string) int {
	days = strings.TrimSpace(days)
	if days == "" {
		log.Fatal("days is empty")
	}

	validDays, err := strconv.Atoi(days)
	if err != nil {
		log.Fatal(err)
	}

	return validDays
}

func NoteValidation(note string) string {
	note = strings.TrimSpace(note)
	if note == "" {
		return ""
	}

	return note
}
