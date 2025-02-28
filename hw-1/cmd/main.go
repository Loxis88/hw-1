package main

import (
	"hw-1/app"
	"hw-1/services"
	"hw-1/storage"
	_ "hw-1/storage/json_storage"
)

func main() {
    storage, err := storage.New("json-storage", "data.json")
    if err != nil {
        panic(err)
    }

    service := services.New(storage)
    app.Serve(service)
}
