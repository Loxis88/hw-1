package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

// HandleReturnOrder processes the return-order command
func HandleReturnOrder(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("return-order", flag.ExitOnError)
	var orderID = flagSet.Uint("order-id", 0, "orderID")
	flagSet.Parse(os.Args[2:])

	if flagSet.NFlag() != 1 || *orderID == 0 {
		fmt.Println("Invalid arguments")
		return
	}

	if err := service.ReturnOrderToCourier(*orderID); err != nil {
		fmt.Println("Error returning order:", err)
	}

	fmt.Println("Order returned successfully")
}
