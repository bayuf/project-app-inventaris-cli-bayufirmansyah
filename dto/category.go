package dto

type GetItemByCategoryDTO struct {
	ID int
}

type CreateCategoryDTO struct {
	Name string
}

type UpdateCategoryDTO struct {
	ID   int
	Name string
}

type CategoryResponseDTO struct {
	ID          int
	Name        string
	Description string
}
