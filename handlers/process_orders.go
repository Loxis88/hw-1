package handlers

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"hw-1/services"
)

const (
	ReturnAction = "return"
	IssueAction  = "issue"
)

func HandleProcessOrders(service services.OrderServiceInterface) error {
	flagSet := flag.NewFlagSet("process-orders", flag.ContinueOnError)

	clientID := flagSet.Uint("client-id", 0, "clientID")
	orderIDs := flagSet.String("order-ids", "", "orderIDs")
	action := flagSet.String("action", "", "action")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	if flagSet.NFlag() < 3 {
		return fmt.Errorf("Invalid arguments")
	}
	if *action != ReturnAction && *action != IssueAction {
		return fmt.Errorf("Invalid action")
	}

	orders := strings.Split(*orderIDs, ",")
	var ids []uint = make([]uint, len(orders))

	for i := range orders {
		id, err := strconv.Atoi(orders[i])
		if err != nil {
			return fmt.Errorf("Invalid order ID:", orders[i])
		}
		ids[i] = uint(id)
	}
	switch *action {
	case "return":
		if err := service.AcceptReturns(*clientID, ids...); err != nil {

			return fmt.Errorf("Error returning orders:", err)
		}
		fmt.Println("Заказы успешно возвращены")
	case "issue":
		if err := service.DeliverOrders(*clientID, ids...); err != nil {
			return fmt.Errorf("Error issueing orders:", err)
		}
		fmt.Println("Заказы успешно выданы")
	}
	return nil
}
