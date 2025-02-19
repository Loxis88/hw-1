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
		fmt.Printf("%v", err)
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
