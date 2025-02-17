package handlers

import (
	"fmt"
	"time"

	"hw-1/services"
)

// HandleAcceptOrder processes the accept-order command
func HandleAcceptOrder(service services.OrderServiceInterface) {
	flagSet := NewFlagSet()
	orderID := flagSet.Uint("order-id", 0, "orderID")
	receiverID := flagSet.Uint("receiver-id", 0, "receiverID")
	storageDuration := flagSet.Uint("storage-duration", 0, "duration")

	flagSet.Parse()

	if flagSet.NFlag() != 3 || *orderID == 0 || *receiverID == 0 || *storageDuration == 0 {
		fmt.Println("Invalid arguments")
		return
	}

	if err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*time.Hour*24)); err != nil {
		fmt.Println("Error accepting order:", err)
	}

	fmt.Println("Order accepted successfully")
}
