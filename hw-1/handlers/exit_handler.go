package handlers

import (
	"fmt"
	"hw-1/cmd/commands"
	"hw-1/services"
	"os"
)

func init() {
	commands.RegisterCommand("exit", commands.Command{
		Description: "Выйти из программы",
		Handle:      HandleExit,
	})
}

// HandleExit завершает выполнение программы.
func HandleExit(service services.OrderServiceInterface) error {
	fmt.Println("Выход из программы...")
	os.Exit(0)
	return nil
}
