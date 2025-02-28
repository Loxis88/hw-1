package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleListReturns(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("list-returns", flag.ContinueOnError)

	page := flagSet.Int("page", 1, "page number (must be positive, default: 1)")
	perPage := flagSet.Int("per-page", 10, "items per page (must be positive, default: 10)")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if *page < 1 {
		return fmt.Errorf("page must be a positive number, got %d", *page)
	}
	if *perPage < 1 {
		return fmt.Errorf("per-page must be a positive number, got %d", *perPage)
	}

	returns, err := service.GetReturnedOrders(*page, *perPage)
	if err != nil {
		return fmt.Errorf("error listing returns: %w", err)
	}

	if len(returns) == 0 {
		fmt.Println("No returns found")
		return nil
	}

	fmt.Println("Returns:")
	for _, order := range returns {
		fmt.Print(order)
	}
	return nil
}
