package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleOrderHistory(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("order-history", flag.ExitOnError)

	limit := flagSet.Int("limit", 0, "limit")
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if *limit < 0 {
		fmt.Println("Invalid arguments: --limit must be greater than or equal to 0")
		return
	}

	history, err := service.GetOrderHistory(*limit)
	if err != nil {
		fmt.Println("Error listing order history:", err)
		return
	}

	fmt.Println("Order History:")
	for _, order := range history {
		fmt.Print(order)
	}
}
