package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"hw-1/services"
)

// HandleProcessOrders processes the process-orders command
func HandleProcessOrders(service services.OrderServiceInterface) {
	flagSet := NewFlagSet()
	clientID := flagSet.Uint("client-id", 0, "clientID")
	orderIDs := flagSet.String("order-ids", "", "orderIDs")
	action := flagSet.String("action", "", "action")
	flagSet.Parse()

	if flagSet.NFlag() < 3 {
		fmt.Println("Invalid arguments")
		return
	}
	if *action != "return" && *action != "issue" {
		fmt.Println("Invalid action")
		return
	}

	orders := strings.Split(*orderIDs, ",")
	var ids []uint = make([]uint, len(orders))

	for i, _ := range orders {
		id, err := strconv.Atoi(orders[i])
		if err != nil {
			fmt.Println("Invalid order ID:", orders[i])
			return
		}
		ids[i] = uint(id)
	}
	switch *action {
	case "return":
		if err := service.AcceptReturns(*clientID, ids...); err != nil {
			fmt.Println("Error accepting orders:", err)
			return
		}
		fmt.Println("Заказы успешно приняты")
		return
	case "issue":
		if err := service.DeliverOrders(*clientID, ids...); err != nil {
			fmt.Println("Error returning orders:", err)
			return
		}
		fmt.Println("Заказы успешно выданы")
		return
	}
}
