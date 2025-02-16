package main

import (
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
}
