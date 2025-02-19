package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleListOrders(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("list-orders", flag.ExitOnError)

	customerID := flagSet.Uint("client-id", 0, "clientID")
	limit := flagSet.Int("limit", 0, "limit")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if *customerID == 0 {
		fmt.Println("Invalid client ID")
		return
	}

	orders, err := service.GetCustomerOrders(*customerID, *limit)
	if err != nil {
		fmt.Println("Error listing orders:", err)
		return
	}

	fmt.Println("Orders:")
	for _, order := range orders {
		fmt.Print(order)
	}
}
