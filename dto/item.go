package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateItemDTO struct {
	Name       string
	CategoryID int
}

type UpdateItemDTO struct {
	ID         int
	Name       *string
	CategoryID *int
}

type GetItemDTO struct {
	CategoryID int
	ItemID     int
}

type ItemResponseDTO struct {
	ID        int
	Name      string
	Price     decimal.Decimal
	BuyDate   time.Time
	TotalUsed int
}
