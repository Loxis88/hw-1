package main

import (
	_ "hw-1/handlers"
	_ "hw-1/storage/json_storage" // нужно для инициализации пакета для выполнения init функции котрая добавляет json-storage в список доступных хранилищ

	"hw-1/cmd/commands"
	"hw-1/services"
	"hw-1/storage"
)

func main() {
	// пока напрямую через указание имени но в будующем можно изменить функцию для работы с конфигом
	storage, err := storage.New("json-storage", "data.json")
	if err != nil {
		panic(err)
	}

	service := services.New(storage)

	commands.Serve(service)
}
