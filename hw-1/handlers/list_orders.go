package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func init() {
	commands.RegisterCommand("list-orders", commands.Command{
		Description: "Получить список заказов\n  Использование: list-orders --client-id <ID> [--limit <N>]",
		Handle:      HandleListOrders,
	})
}

func HandleListOrders(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("list-orders", flag.ContinueOnError)

	customerID := flagSet.Uint("client-id", 0, "clientID")
	limit := flagSet.Int("limit", 0, "limit")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if *customerID == 0 {
		return fmt.Errorf("Флаг client-id не указан")
	}

	orders, err := service.GetCustomerOrders(*customerID, *limit)
	if err != nil {
		return fmt.Errorf("Error listing orders: %w", err)
	}

	fmt.Println("Orders:")
	for _, order := range orders {
		fmt.Print(order)
	}
	return nil
}
