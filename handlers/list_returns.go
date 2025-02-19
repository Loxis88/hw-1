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
		fmt.Printf("%v", err)
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
