package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hw-1/cmd/commands"
	"hw-1/services"
	"hw-1/storage"
)

func main() {
	store, err := storage.NewJsonStorage("data.json")
	if err != nil {
		panic(err)
	}

	var service services.OrderServiceInterface = services.New(store)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		commandName := os.Args[1]
		cmd, exists := commands.Commands[commandName]
		if !exists {
			fmt.Println("Invalid command:", commandName)
			commands.Commands["help"].Handle(service) // Вывод справки при ошибке
			return
		}

		// Выполняем обработчик команды
		err := cmd.Handle(service)
	}
}
