package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

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
