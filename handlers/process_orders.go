package handlers

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"hw-1/cmd/commands"
	"hw-1/services"
)

const (
	ReturnAction = "return"
	IssueAction  = "issue"
)

func init() {
	commands.RegisterCommand("process-orders", commands.Command{
		Description: "Выдать заказы или принять возвраты\n  Использование: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <issue|return>",
		Handle:      HandleProcessOrders,
	})
}

func HandleProcessOrders(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("process-orders", flag.ContinueOnError)

	clientID := flagSet.Uint("client-id", 0, "client ID (required, must be positive)")
	orderIDs := flagSet.String("order-ids", "", "comma-separated order IDs (required)")
	action := flagSet.String("action", "", "action: 'return' or 'issue' (required)")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	missingFlags := []string{}
	if *clientID == 0 {
		missingFlags = append(missingFlags, "-client-id")
	}
	if *orderIDs == "" {
		missingFlags = append(missingFlags, "-order-ids")
	}
	if *action == "" {
		missingFlags = append(missingFlags, "-action")
	}

	if len(missingFlags) > 0 {
		return fmt.Errorf("missing required flags: %s", strings.Join(missingFlags, ", "))
	}

	if *clientID == 0 {
		return fmt.Errorf("client-id must be a positive number, got %d", *clientID)
	}

	if *orderIDs == "" {
		return fmt.Errorf("order-ids cannot be empty")
	}
	orderStrs := strings.Split(*orderIDs, ",")
	if len(orderStrs) == 0 {
		return fmt.Errorf("order-ids must contain at least one ID")
	}

	ids := make([]uint, 0, len(orderStrs))
	for _, orderStr := range orderStrs {
		orderStr = strings.TrimSpace(orderStr)
		id, err := strconv.ParseUint(orderStr, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid order ID '%s': %w", orderStr, err)
		}
		ids = append(ids, uint(id))
	}

	switch *action {
	case ReturnAction:
		if err := service.AcceptReturns(*clientID, ids...); err != nil {
			return fmt.Errorf("error returning orders: %w", err)
		}
		fmt.Println("Orders returned successfully")
	case IssueAction:
		if err := service.IssueOrders(*clientID, ids...); err != nil {
			return fmt.Errorf("error issuing orders: %w", err)
		}
		fmt.Println("Orders issued successfully")
	default:
		return fmt.Errorf("invalid action '%s', must be '%s' or '%s'", *action, ReturnAction, IssueAction)
	}

	return nil
}
