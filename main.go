package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price int `json:"price"`
}
func loadMenu() error{
	data, err := os.ReadFile("data/menu.json")
	if err != nil {
		return  err
	}

	err = json.Unmarshal(data, &menus)
	if err != nil {
		return err
	}
	return  nil
}
var menus =[]Menu{}



func showMenu() {
	
	fmt.Println("=============================================================")
	fmt.Println("||              ESTEH INDONESIA NUSANTARA                  ||")
	fmt.Println("=============================================================")
	fmt.Printf("%-3s %-30s %-20s %-10s\n", "ID", "MENU", "CATEGORY", "PRICE")
	fmt.Println("--------------------------------------------------------------")
	for _, menu := range menus{
	fmt.Printf("%-3d %-30s %-20s Rp.%-10d\n",menu.ID,menu.Name,menu.Category,menu.Price)
	}
}

func showMainMenu()string{

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=============================================================")
	fmt.Println("||              ESTEH INDONESIA NUSANTARA                  ||")
	fmt.Println("=============================================================")
	fmt.Println("")
	fmt.Println("1.Show Menu")
	fmt.Println("2.Add to cart")
	fmt.Println("3.View Cart")
	fmt.Println("4.Checkout")
	fmt.Println("")
	fmt.Println("0.Exit")
	fmt.Println("")

	fmt.Print("Choose the menu :")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return ""
	}
	input = strings.TrimSpace(input)
	return  input
}

func findMenuByID(id int) *Menu {
	for index, menu := range menus{
		if menu.ID == id {
			return  &menus[index]
		}
	}
		return nil
}

func main() {
	err := loadMenu()
	if err != nil {
		fmt.Println(err)
		return 
	}

	for {
		choice := showMainMenu()

	switch choice {
	case "1":
    	showMenu()
	case "0":
    	fmt.Println("Thank you")
    return
	default:
    	fmt.Println("Menu tidak tersedia")
	}
	}
	
}