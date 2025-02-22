package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func init() {
	commands.RegisterCommand("order-history", commands.Command{
		Description: "Получить историю заказов\n  Использование: order-history [--limit]",
		Handle:      HandleOrderHistory,
	})
}

func HandleOrderHistory(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("order-history", flag.ContinueOnError)
	limit := flagSet.Int("limit", 0, "limit")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %w\n", err)
	}

	if *limit < 0 {
		return fmt.Errorf("Invalid arguments: --limit must be greater than or equal to 0")
	}

	history, err := service.GetOrderHistory(*limit)
	if err != nil {
		return fmt.Errorf("Error listing order history: %w", err)
	}

	fmt.Println("Order History:")
	for _, order := range history {
		fmt.Print(order)
	}
	return nil
}
