package main

import (
	"fmt"
	"hw-1/services"
	"hw-1/storage"
	"time"
)

func main() {
	store, err := storage.NewJsonStorage("data.json")
	if err != nil {
		panic(err)
	}

	var service services.OrderServiceInterface = services.New(store)

	service.AcceptOrder(10, 15, time.Now().Add(time.Hour))

	// Call GetOrderHistory method and print the result
	orderHistory, err := service.GetOrderHistory(10)
	if err != nil {
		fmt.Println("Error fetching order history:", err)
	} else {
		fmt.Println("Order History:", orderHistory)
	}

	// Call GetReturnedOrders method and print the result
	returnedOrders, err := service.GetReturnedOrders()
	if err != nil {
		fmt.Println("Error fetching returned orders:", err)
	} else {
		fmt.Println("Returned Orders:", returnedOrders)
	}
}
