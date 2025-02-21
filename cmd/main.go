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

		args := strings.Fields(input)
		commandName := args[0]

		// замена os.Args на инпут
		os.Args = args
		// тут кароче костыль потому что я не могу полоижть команду помощи в мапу
		if commandName == commands.HelpCommand {
			commands.HandleHelp()
			continue
		}

		cmd, exists := commands.Commands[commandName]
		if !exists {
			fmt.Println("Invalid command:", commandName)
			commands.HandleHelp()
			continue
		}

		// Выполняем обработчик команды
		err = cmd.Handle(service)
		if err != nil {
			fmt.Println(err)
		}
	}
}
