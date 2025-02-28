package handlers

import (
	"fmt"
	"sort"
	"strings"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func HandleHelp(service services.OrderServiceInterface) error {
	registeredCommands := commands.GetRegisteredCommands()

	fmt.Println("Доступные команды:")

	commandNames := make([]string, 0, len(registeredCommands))
	for name := range registeredCommands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	for _, name := range commandNames {
		cmd := registeredCommands[name]
		description := strings.ReplaceAll(cmd.Description, "\n  ", "\n    ")
		fmt.Printf("  %s\n    %s\n", name, description)
	}

	return nil
}
