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
	ID     int
	Name   string
	Note   string
	Status string
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
	Status     string
	Note       string
	LifeDays   int

	TotalInvestment   decimal.Decimal
	CurrentValue      decimal.Decimal
	DepreciationValue decimal.Decimal
	AgeItem           float64
}
