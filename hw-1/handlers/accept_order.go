package handlers

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"hw-1/models"
	"hw-1/services"
)

func HandleAcceptOrder(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("accept-order", flag.ContinueOnError)

	orderID := flagSet.Uint("order-id", 0, "order ID (required, must be positive)")
	receiverID := flagSet.Uint("receiver-id", 0, "receiver ID (required, must be positive)")
	storageDuration := flagSet.Uint("storage-duration", 0, "storage duration in days (required, must be positive)")
	cost := flagSet.Float64("cost", 0, "order cost (required)")
	weight := flagSet.Float64("weight", 0, "order weight (required)")
	packageType := flagSet.String("package", "", "order package (optional)")
	wrap := flagSet.Bool("add-wrap", false, "add film warp (optional) <true/false>")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	missingFlags := []string{}
	if *orderID == 0 {
		missingFlags = append(missingFlags, "-order-id")
	}
	if *receiverID == 0 {
		missingFlags = append(missingFlags, "-receiver-id")
	}
	if *storageDuration == 0 {
		missingFlags = append(missingFlags, "-torage-duration")
	}
	if *cost == 0 {
		missingFlags = append(missingFlags, "-cost")
	}
	if *weight == 0 {
		missingFlags = append(missingFlags, "-weight")
	}

	if len(missingFlags) > 0 {
		return fmt.Errorf("missing required flags: %s", strings.Join(missingFlags, ", "))
	}

	err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*24*time.Hour), *weight, *cost, models.PackageType(*packageType), *wrap)
	if err != nil {
		return fmt.Errorf("error accepting order: %w", err)
	}

	fmt.Println("Order accepted successfully")
	return nil
}
