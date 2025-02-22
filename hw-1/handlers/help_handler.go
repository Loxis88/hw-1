package handlers

import (
	"fmt"
	"os"
	"slices"

	"hw-1/cmd/commands"
	"hw-1/services"
)

func init() {
	commands.RegisterCommand("help", commands.Command{
		Description: "Ввести инормацию по командам (команде)\n  Использование: help [имя команды] (необязательный параметр)\n",
		Handle:      HandleHelp,
	})
}

func HandleHelp(service services.OrderServiceInterface) error {
	if len(os.Args) > 2 {
		return fmt.Errorf("too many arguments")
	}
	if len(os.Args) == 2 {
		if command, ok := commands.Commands[os.Args[1]]; ok {
			fmt.Printf("%s\n  %s\n", os.Args[1], command.Description)
		} else {
			return fmt.Errorf("command %s not found\n", os.Args[1])
		}
		return nil
	}

	fmt.Println("Доступные команды:")
	cmd := []string{}

	for name, _ := range commands.Commands {
		cmd = append(cmd, name)
	}

	slices.Sort(cmd)
	for _, name := range cmd {
		fmt.Printf("%s\n  %s\n", name, commands.Commands[name].Description)
	}
	return nil
}
