package handlers

import (
	"flag"
	"fmt"
	"time"

	"hw-1/services"
)

// HandleAcceptOrder processes the accept-order command
func HandleAcceptOrder(service services.OrderServiceInterface) {
	orderID := flag.Uint("order-id", 0, "orderID")
	receiverID := flag.Uint("receiver-id", 0, "receiverID")
	storageDuration := flag.Uint("storage-duration", 0, "duration")

	flag.Parse()

	if flag.NFlag() != 3 || *orderID == 0 || *receiverID == 0 || *storageDuration == 0 {
		fmt.Println("Invalid arguments")
		return
	}

	if err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*time.Hour*24)); err != nil {
		fmt.Println("Error accepting order:", err)
	}

	fmt.Println("Order accepted successfully")
}
