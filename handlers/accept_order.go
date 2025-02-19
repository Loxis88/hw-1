package handlers

import (
	"flag"
	"fmt"
	"os"
	"time"

	"hw-1/services"
)

func HandleAcceptOrder(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("accept-order", flag.ContinueOnError)

	orderID := flagSet.Uint("order-id", 0, "orderID")
	receiverID := flagSet.Uint("receiver-id", 0, "receiverID")
	storageDuration := flagSet.Uint("storage-duration", 0, "duration")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if flagSet.NFlag() != 3 || *orderID == 0 || *receiverID == 0 || *storageDuration == 0 {
		fmt.Println("Invalid arguments", *orderID, *receiverID, *storageDuration)
		return
	}

	if err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*time.Hour*24)); err != nil {
		fmt.Println("Error accepting order:", err)
		return
	}

	fmt.Println("Orders accepted successfully")
}
