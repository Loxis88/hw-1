package handlers

import (
	"flag"
	"fmt"

	"hw-1/services"
)

// HandleReturnOrder processes the return-order command
func HandleReturnOrder(service services.OrderServiceInterface) {
	var orderID = flag.Uint("order-id", 0, "orderID")
	flag.Parse()

	if flag.NFlag() != 1 || *orderID == 0 {
		fmt.Println("Invalid arguments")
		return
	}

	if err := service.ReturnOrderToCourier(*orderID); err != nil {
		fmt.Println("Error returning order:", err)
	}

	fmt.Println("Order returned successfully")
}
