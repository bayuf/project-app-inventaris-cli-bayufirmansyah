package model

type ItemCategory struct {
	ID          int
	Name        string
	Description string
	Quantity    int
}

type Item struct {
	ID   int
	Name string
}
