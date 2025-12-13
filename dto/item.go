package dto

type CreateItemDTO struct {
	Name       string
	CategoryID int
	Quantity   int
}

type UpdateItemDTO struct {
	ID         int
	Name       *string
	CategoryID *int
	Quantity   *int
}

type GetItemDTO struct {
	CategoryID int
	ItemID     int
}

type ItemResponseDTO struct {
	ID         int
	Name       string
	CategoryID int
	Quantity   int
	CreateAt   string
	UpdateAt   string
}
