package model

type Order struct {
	ID     int
	UserID int
	Items  []Item
}

type Item struct {
	ID          int
	Name        string
	MainShelfID int
	Quantity    int
	Shelves     []Shelf
}

type Shelf struct {
	ID   int
	Name string
}
