package cart

import "koda-b8-Weekly3/menu"


type CartItem struct {
    Menu menu.Menu
    Quantity int
}

var cart []CartItem