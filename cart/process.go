package cart

import (
	"fmt"
	"sync"
	"time"
)

var (
	logs  []string
	mutex sync.Mutex
)

func SaveTransaction(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Saving transaction...")
	time.Sleep(4 * time.Second)

	mutex.Lock()
	defer mutex.Unlock()

	logs = append(logs, "Transaction saved")
}

func UpdateStock(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Updating stock...")
	time.Sleep(3 * time.Second)
	mutex.Lock()
	defer mutex.Unlock()

	logs = append(logs, "Stock updated")
}

func PrintReceiptProcess(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Printing receipt...")
	time.Sleep(5 * time.Second)
	mutex.Lock()
	defer mutex.Unlock()

	logs = append(logs, "Receipt printed")
}

func ShowLogs() {
	fmt.Println("\n===== Checkout Process =====")

	for _, log := range logs {
		fmt.Println(log)
	}

	logs = nil
}