package main

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
func main() {
	err := loadMenu()
	if err != nil {
		fmt.Println(err)
		return 
	}
	// fmt.Printf(`id : %d, name : %s, price:%d`,menus.ID)
	showMenu()

}