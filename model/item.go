package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ItemCategory struct {
	ID          int
	Name        string
	Description string
	Quantity    int
}

type Item struct {
	ID        int
	Name      string
	Price     decimal.Decimal
	BuyDate   time.Time
	TotalUsed int
}
