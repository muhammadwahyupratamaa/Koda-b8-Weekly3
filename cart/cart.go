package cart

import (
	"fmt"
	"koda-b8-Weekly3/menu"
)


type CartItem struct {
    Menu menu.Menu
    Quantity int
}

var cart []CartItem

func AddToCart(menuItem menu.Menu, qty int) {
	for i := range cart {
		if cart[i].Menu.ID == menuItem.ID {
			cart[i].Quantity += qty
			fmt.Println("Quantity updated successfully!")
			return
		}
	}

	cart = append(cart, CartItem{
		Menu:     menuItem,
		Quantity: qty,
	})

	fmt.Println("Menu successfully added to cart!")
}

func ViewCart() {
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