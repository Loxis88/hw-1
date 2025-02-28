package handlers

import (
	"fmt"
	"hw-1/services"
	"os"
)

// HandleExit завершает выполнение программы.
func HandleExit(service services.OrderServiceInterface) error {
	fmt.Println("Выход из программы...")
	os.Exit(0)
	return nil
}
