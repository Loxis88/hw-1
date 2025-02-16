package main

import (
	"hw-1/services"
	"hw-1/storage/json_storage"
	"time"
)

func main() {
	store, err := json_storage.New("data.json")
	if err != nil {
		panic(err)
	}

	service := services.New(store)
	service.AcceptOrder(3, 3, time.Now().Add(time.Hour))
}
