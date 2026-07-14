package menu

type Repository interface {
	LoadMenu() error
	FindMenuByID(id int) *Menu
	GetCategories() []string
	ShowMenuByCategory(category string)
}