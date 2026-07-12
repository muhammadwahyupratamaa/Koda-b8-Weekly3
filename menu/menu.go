package menu

import (
	"encoding/json"
	"fmt"
	"os"
)

type Menu struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price int `json:"price"`
}

var Menus []Menu

func LoadMenu() error{
	data, err := os.ReadFile("data/menu.json")
	if err != nil {
		return  err
	}

	err = json.Unmarshal(data, &Menus)
	if err != nil {
		return err
	}
	return  nil
}

func ShowMenu() {
	fmt.Println("\n=== Menu List ===")

	for _, menu := range Menus {
		fmt.Println("----------------------------")
		fmt.Println("ID       :", menu.ID)
		fmt.Println("Menu     :", menu.Name)
		fmt.Println("Category :", menu.Category)
		fmt.Println("Price    : Rp.", menu.Price)
	}

	fmt.Println("----------------------------")
}

func FindMenuByID(id int) *Menu {
	for index, menu := range Menus{
		if menu.ID == id {
			return  &Menus[index]
		}
	}
		return nil
}

func GetCategories() []string {
	categories := []string{}
	visited := make(map[string]bool)

	for _, menu := range Menus {
		if !visited[menu.Category] {
			categories = append(categories, menu.Category)
			visited[menu.Category] = true
		}
	}

	return categories
}

func ShowMenuByCategory(category string) {
	fmt.Printf("\n=== %s ===\n", category)

	for _, menu := range Menus {
		if menu.Category == category {
			fmt.Println("----------------------------")
			fmt.Println("ID       :", menu.ID)
			fmt.Println("Menu     :", menu.Name)
			fmt.Println("Category :", menu.Category)
			fmt.Println("Price    : Rp.", menu.Price)
		}
	}

	fmt.Println("----------------------------")
}