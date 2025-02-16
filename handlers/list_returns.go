package handlers

import (
	"fmt"

	"hw-1/services"
)

// HandleListReturns processes the list-returns command
func HandleListReturns(service services.OrderServiceInterface) {
	returns, err := service.GetReturnedOrders()
	if err != nil {
		fmt.Println("Error listing returns:", err)
		return
	}
	fmt.Println("Returns:", returns)
}
