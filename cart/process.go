package cart

import (
	"fmt"
	"sync"
	"time"
)

func SaveTransaction(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Saving transaction...")
	time.Sleep(4 * time.Second)
	fmt.Println("Transaction saved")
}

func UpdateStock(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Updating stock...")
	time.Sleep(3 * time.Second)
	fmt.Println("Stock updated")
}

func PrintReceiptProcess(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Printing receipt...")
	time.Sleep(5 * time.Second)
	fmt.Println("Receipt printed")
}