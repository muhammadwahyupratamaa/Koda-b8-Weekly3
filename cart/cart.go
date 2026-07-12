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