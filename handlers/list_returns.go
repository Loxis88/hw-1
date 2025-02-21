package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleListReturns(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("list-returns", flag.ContinueOnError)

	page := flagSet.Int("page", 1, "page")
	perPage := flagSet.Int("per-page", 10, "per-page")

	if err := flagSet.Parse(os.Args[1:]); err != nil {

		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if *page < 1 || *perPage < 1 {
		return fmt.Errorf("Invalid arguments: --page and --per-page must be greater than 0")
	}

	returns, err := service.GetReturnedOrders(*page, *perPage)
	if err != nil {
		return fmt.Errorf("Error listing returns:", err)
	}
	fmt.Println("Returns")
	for _, order := range returns {
		fmt.Print(order)
	}
	return nil
}
