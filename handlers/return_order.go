package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func init() {
	commands.RegisterCommand("return-order", commands.Command{
		Description: "Вернуть заказ курьеру\n  Использование: return-order --order-id <ID>",
		Handle:      HandleReturnOrder,
	})
}

func HandleReturnOrder(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("return-order", flag.ContinueOnError)
	var orderID = flagSet.Uint("order-id", 0, "orderID")
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if flagSet.NFlag() != 1 || *orderID == 0 {
		return fmt.Errorf("Invalid arguments: --order-id is required")
	}

	if err := service.ReturnOrderToCourier(*orderID); err != nil {
		return fmt.Errorf("Error returning order: %w", err)
	}

	fmt.Println("Order returned successfully")
	return nil
}
