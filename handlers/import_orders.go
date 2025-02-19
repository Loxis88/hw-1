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
		fmt.Printf("%v", err)
		return
	}

	if *path == "" {
		return
	}

	if err := service.ImportOrders(*path); err != nil {
		fmt.Printf("Ошибка при импорте: %v\n", err)
		return
	}

	fmt.Println("Заказы успешно импортированы")
}
