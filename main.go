package main

import (
	"bufio"
	"fmt"
	"koda-b8-Weekly3/cart"
	"koda-b8-Weekly3/menu"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var repo menu.Repository = menu.JSONRepository{}

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
var reader = bufio.NewReader(os.Stdin)
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

	menuItem := repo.FindMenuByID(id)
	if menuItem == nil {
		fmt.Println("Menu not found!")
		return
	}

	if menuItem.Category != category {
		fmt.Println("Menu is not in this category!")
		return
	}

	qty, err := inputQuantity()
	if err != nil {
		fmt.Println(err)
		return
	}

	cart.AddToCart(*menuItem, qty)
}

func checkout() {
	if cart.IsEmpty() {
		fmt.Println("Cart is empty!")
		return
	}
	cart.ViewCart()
	total := cart.TotalPayment()
	for {
		payment := inputPayment()

		if payment < total {
			fmt.Println("Your money is not enough!")
			continue
		}

		change := cart.CalculateChange(payment, total)
		go cart.SaveTransaction()
		go cart.UpdateStock()
		go cart.PrintReceiptProcess()
		cart.PrintReceipt(total, payment, change)
		cart.Clear()

		break
	}
}

func inputPayment() int {
	for {
		fmt.Print("Input Payment : Rp. ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSpace(input)

		payment, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Payment must be a number!")
			continue
		}

		if payment <= 0 {
			fmt.Println("Payment must be greater than 0!")
			continue
		}

		return payment
	}
}

func showCategory() {
	categories := repo.GetCategories()

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
	repo.ShowMenuByCategory(selectedCategory)
	processAddToCart(selectedCategory)
}



func main() {
	ClearScreen()
	err := repo.LoadMenu()
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
		cart.ViewCart()
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