package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Menu struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price int `json:"price"`
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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
	fmt.Println("\n=== Menu List ===")

	for _, menu := range menus {
		fmt.Println("----------------------------")
		fmt.Println("ID       :", menu.ID)
		fmt.Println("Menu     :", menu.Name)
		fmt.Println("Category :", menu.Category)
		fmt.Println("Price    : Rp.", menu.Price)
	}

	fmt.Println("----------------------------")
}

func showMainMenu()string{
	fmt.Println("")
	fmt.Println("=============================================================")
	fmt.Println("||              ESTEH INDONESIA NUSANTARA                  ||")
	fmt.Println("=============================================================")
	fmt.Println("")
	fmt.Println("1.Show Menu")
	fmt.Println("2.View Cart")
	fmt.Println("3.Checkout")
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

func inputMenuID() (int, error) {
	fmt.Print("Input Menu ID : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("menu ID must be a number")
	}

	return id, nil
}

type CartItem struct {
	Menu     Menu
	Quantity int
}

var cart []CartItem

func addToCart(menu Menu, qty int) {
	for i := range cart {
		if cart[i].Menu.ID == menu.ID {
			cart[i].Quantity += qty
			fmt.Println("Quantity updated successfully!")
			return
		}
	}

	cart = append(cart, CartItem{
		Menu:     menu,
		Quantity: qty,
	})

	fmt.Println("Menu successfully added to cart!")
}

func inputQuantity() (int, error) {
	fmt.Print("Input Quantity : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	qty, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("quantity must be a number")
	}

	if qty <= 0 {
		return 0, fmt.Errorf("quantity must be greater than 0")
	}

	return qty, nil
}

func processAddToCart(category string) {
	id, err := inputMenuID()
	if err != nil {
		fmt.Println(err)
		return
	}

	menu := findMenuByID(id)
	if menu == nil {
		fmt.Println("Menu not found!")
		return
	}

	if menu.Category != category {
		fmt.Println("Menu is not in this category!")
		return
	}

	qty, err := inputQuantity()
	if err != nil {
		fmt.Println(err)
		return
	}

	addToCart(*menu, qty)
}

func viewCart() {
	if len(cart) == 0 {
		fmt.Println("Cart is empty!")
		return
	}

	fmt.Println("\n=== Your Cart ===")

	for _, item := range cart {
		fmt.Println("----------------------------")
		fmt.Println("ID       :", item.Menu.ID)
		fmt.Println("Menu     :", item.Menu.Name)
		fmt.Println("Category :", item.Menu.Category)
		fmt.Println("Price    : Rp.", item.Menu.Price)
		fmt.Println("Quantity :", item.Quantity)
		fmt.Println("Subtotal : Rp.", item.Menu.Price*item.Quantity)
	}

	fmt.Println("----------------------------")
	fmt.Println("Total Payment : Rp.", totalPayment())
}

func inputPayment() (int, error) {
	fmt.Print("Input Payment : Rp. ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	payment, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("payment must be a number")
	}

	if payment <= 0 {
		return 0, fmt.Errorf("payment must be greater than 0")
	}

	return payment, nil
}

func totalPayment() int {
	total := 0

	for _, item := range cart {
		total += item.Menu.Price * item.Quantity
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

func getCategories() []string {
	categories := []string{}
	visited := make(map[string]bool)

	for _, menu := range menus {
		if !visited[menu.Category] {
			categories = append(categories, menu.Category)
			visited[menu.Category] = true
		}
	}

	return categories
}

func showCategory() {
	categories := getCategories()

	fmt.Println("\n=== Category ===")

	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category)
	}

	fmt.Print("\nChoose Category : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input must be a number!")
		return
	}

	if choice < 1 || choice > len(categories) {
		fmt.Println("Category not found!")
		return
	}

	selectedCategory := categories[choice-1]
	ClearScreen()
	showMenuByCategory(selectedCategory)
	processAddToCart(selectedCategory)
}

func showMenuByCategory(category string) {
	fmt.Printf("\n=== %s ===\n", category)

	for _, menu := range menus {
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

func main() {
	ClearScreen()
	err := loadMenu()
	if err != nil {
		fmt.Println(err)
		return 
	}

	for {
		choice := showMainMenu()

	switch choice {
	case "1":
		ClearScreen()
		showCategory()
	case "2":
		viewCart()
	case "3":
		checkout()
	case "0":
		fmt.Println("Thank you")
		return
	default:
		fmt.Println("Menu tidak tersedia")
	}
	}
	
}