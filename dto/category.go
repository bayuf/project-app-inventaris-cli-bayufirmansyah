package dto

type GetItemByCategoryDTO struct {
	ID int
}

type CreateCategoryDTO struct {
	Name        string
	Description string
}

type UpdateCategoryDTO struct {
	ID          int
	Name        string
	Description string
}

type CategoryResponseDTO struct {
	ID          int
	Name        string
	Description string
	Quantity    int
}
