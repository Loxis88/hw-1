package handlers

import (
	"flag"
	"fmt"

	"hw-1/services"
)

// HandleListOrders processes the list-orders command
func HandleListOrders(service services.OrderServiceInterface) {
	customerID := flag.Uint("client-id", 0, "clientID")
	limit := flag.Int("limit", 0, "limit")
	flag.Parse()

	if *customerID == 0 {
		fmt.Println("Invalid client ID")
		return
	}

	orders, err := service.GetCustomerOrders(*customerID, *limit)
	if err != nil {
		fmt.Println("Error listing orders:", err)
		return
	}
	fmt.Println("Orders:", orders)
}
