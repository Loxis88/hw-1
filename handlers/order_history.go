package handlers

import (
	"flag"
	"fmt"

	"hw-1/services"
)

// HandleOrderHistory processes the order-history command
func HandleOrderHistory(service services.OrderServiceInterface) {
	customerID := flag.Int("client-id", 0, "clientID")
	flag.Parse()

	if *customerID == 0 {
		fmt.Println("Invalid client ID")
		return
	}

	history, err := service.GetOrderHistory(*customerID)
	if err != nil {
		fmt.Println("Error listing order history:", err)
		return
	}
	fmt.Println("Order History:", history)
}
