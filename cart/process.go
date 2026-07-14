package cart

import (
	"fmt"
	"time"
)

func SaveTransaction() {
	fmt.Println("Saving transaction...")
	time.Sleep(2 * time.Second)
	fmt.Println("Transaction saved")
}

func UpdateStock() {
	fmt.Println("Updating stock...")
	time.Sleep(1 * time.Second)
	fmt.Println("Stock updated")
}

func PrintReceiptProcess() {
	fmt.Println("Printing receipt...")
	time.Sleep(3 * time.Second)
	fmt.Println("Receipt printed")
}