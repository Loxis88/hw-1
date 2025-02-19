package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleImportOrders(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("import", flag.ContinueOnError)

	path := flagSet.String("path", "", "path to json orders")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if *path == "" {
		fmt.Println("Invalid arguments: --path is required")
		return
	}

	if err := service.ImportOrders(*path); err != nil {
		fmt.Printf("Ошибка при импорте: %v\n", err)
		return
	}

	fmt.Println("Заказы успешно импортированы")
}
