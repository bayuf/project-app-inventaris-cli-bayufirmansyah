package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateItemDTO struct {
	ID         int
	Category   string
	Name       string
	Price      decimal.Decimal
	BuyDate    time.Time
	TotalUsage int

	CategoryId int
	SKU        string
	LifeDays   int
	Note       string
}

type UpdateItemDTO struct {
	ID         int
	Name       *string
	CategoryID *int
}

type GetItemDTO struct {
	ItemID int
}

type ItemResponseDTO struct {
	ID         int
	Category   string
	Name       string
	Price      decimal.Decimal
	BuyDate    time.Time
	TotalUsage int

	CategoryId int
	SKU        string
}
