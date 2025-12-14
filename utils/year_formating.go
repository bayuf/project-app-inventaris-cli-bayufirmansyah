package utils

import (
	"fmt"
	"math"
	"time"
)

func YearFormat(year float64) string {

	const hoursInYear = 365.25 * 24

	totalDuration := time.Duration(year*hoursInYear) * time.Hour

	years := int(totalDuration.Hours() / (365 * 24))

	remainingHours := totalDuration.Hours() - float64(years*365*24)
	remainingDuration := time.Duration(remainingHours) * time.Hour
	const hoursInMonth = 30 * 24
	months := int(remainingDuration.Hours() / hoursInMonth)
	remainingDaysHours := remainingDuration.Hours() - float64(months*hoursInMonth)
	days := int(math.Floor(remainingDaysHours / 24))

	return fmt.Sprintf("%d Tahun, %d Bulan, %d Hari", years, months, days)

}
