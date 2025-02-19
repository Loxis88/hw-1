package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleListReturns(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("list-returns", flag.ExitOnError)

	page := flagSet.Int("page", 1, "page")
	perPage := flagSet.Int("per-page", 10, "per-page")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if *page < 1 || *perPage < 1 {
		fmt.Println("Invalid arguments: --page and --per-page must be greater than 0")
		return
	}

	returns, err := service.GetReturnedOrders(*page, *perPage)
	if err != nil {
		fmt.Println("Error listing returns:", err)
		return
	}
	fmt.Println("Returns")
	for _, order := range returns {
		fmt.Print(order)
	}
}
