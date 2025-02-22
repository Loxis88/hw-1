package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func init() {
	commands.RegisterCommand("import", commands.Command{
		Description: "Импортировать заказы\n  Использование: import [--path <путь к json>]",
		Handle:      HandleImportOrders,
	})
}

func HandleImportOrders(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("import", flag.ContinueOnError)

	path := flagSet.String("path", "", "path to json orders")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if *path == "" {
		return fmt.Errorf("Invalid arguments: --path is required")
	}

	if err := service.ImportOrders(*path); err != nil {
		return fmt.Errorf("Ошибка при импорте: %v\n", err)
	}

	fmt.Println("Заказы успешно импортированы")
	return nil
}
