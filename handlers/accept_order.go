package handlers

import (
	"flag"
	"fmt"
	"os"
	"time"

	"hw-1/services"
)

func HandleAcceptOrder(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("accept-order", flag.ContinueOnError)

	orderID := flagSet.Uint("order-id", 0, "orderID")
	receiverID := flagSet.Uint("receiver-id", 0, "receiverID")
	storageDuration := flagSet.Uint("storage-duration", 0, "duration")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if flagSet.NFlag() != 3 {

		return fmt.Errorf("Invalid arguments", *orderID, *receiverID, *storageDuration)
	}

	if err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*time.Hour*24)); err != nil {
		return fmt.Errorf("Error accepting order:", err)
	}

	fmt.Println("Orders accepted successfully")
	return nil
}
