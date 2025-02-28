// hw-1/app/app.go
package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hw-1/cli/registry"
	"hw-1/services"
)

// Serve обрабатывает пользовательский ввод и выполняет команды
func Serve(service services.OrderServiceInterface) {
	cmds := registry.RegisterCommands()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		commandName := args[0]

		os.Args = args

		cmd, exists := cmds[commandName]
		if !exists {
			fmt.Println("Invalid command:", commandName, "type help for more information")
			continue
		}

		err = cmd.Handle(service)
		if err != nil {
			fmt.Println(err)
		}
	}
}
