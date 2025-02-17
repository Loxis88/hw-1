package handlers

import (
	"fmt"

	"hw-1/services"
)

// HandleListOrders processes the list-orders command
func HandleListOrders(service services.OrderServiceInterface) {
	flagSet := NewFlagSet()
	customerID := flagSet.Uint("client-id", 0, "clientID")
	limit := flagSet.Int("limit", 0, "limit")
	flagSet.Parse()

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
