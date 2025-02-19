package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleReturnOrder(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("return-order", flag.ExitOnError)
	var orderID = flagSet.Uint("order-id", 0, "orderID")
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if flagSet.NFlag() != 1 || *orderID == 0 {
		fmt.Println("Invalid arguments: --order-id is required")
		return
	}

	if err := service.ReturnOrderToCourier(*orderID); err != nil {
		fmt.Println("Error returning order:", err)
		return
	}

	fmt.Println("Order returned successfully")
}
