package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
var reader = bufio.NewReader(os.Stdin)
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
	fmt.Println("")
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

var cart []Menu

func addToCart() {
	showMenu()
	fmt.Println("")

	fmt.Print("Input Menu ID : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("The input must be a number!")
		return
	}
	menu := findMenuByID(id)
	if menu == nil {
		fmt.Println("Menu Not Found!")
	}
	cart = append(cart, *menu)

	fmt.Println("Menu successfully added to cart!")
}

func viewCart() {
	if len(cart) == 0 {
		fmt.Println("Cart is Empty!")
		return
	}

	fmt.Println("=============================================================")
	fmt.Println("||                      YOUR CART                         ||")
	fmt.Println("=============================================================")
	fmt.Printf("%-3s %-30s %-20s %-10s\n", "ID", "MENU", "CATEGORY", "PRICE")
	fmt.Println("--------------------------------------------------------------")

	for _, item := range cart {
		fmt.Printf("%-3d %-30s %-20s Rp.%-10d\n",
			item.ID,
			item.Name,
			item.Category,
			item.Price,
		)
	}
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Total : Rp.%d\n", totalPayment())
}

func totalPayment() int {
	total := 0

	for _, item := range cart {
		total += item.Price
	}

	return total
}

func checkout() {
	if len(cart) == 0 {
		fmt.Println("Cart is Empty!")
		return
	}

	viewCart()

	fmt.Println("==============================")
	fmt.Println("Checkout Success!")
	fmt.Printf("Total Payment : Rp.%d\n", totalPayment())
	fmt.Println("==============================")

	cart = nil
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
	case "2":
		addToCart()
	case "3":
		viewCart()
	case "4":
		checkout()
	case "0":
    	fmt.Println("Thank you")
    return
	default:
    	fmt.Println("Menu tidak tersedia")
	}
	}
	
}