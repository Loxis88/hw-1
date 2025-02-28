package handlers

import (
	"fmt"
	"os"
	"slices"

	"hw-1/cmd/commands"
	"hw-1/services"
)

// HandleHelp показывает информацию о доступных командах
func HandleHelp(service services.OrderServiceInterface, commandsList map[string]commands.Command) error {
    if len(os.Args) > 2 {
        return fmt.Errorf("too many arguments")
    }

    if len(os.Args) == 2 {
        if command, ok := commandsList[os.Args[1]]; ok {
            fmt.Printf("%s\n  %s\n", os.Args[1], command.Description)
        } else {
            return fmt.Errorf("command %s not found\n", os.Args[1])
        }
        return nil
    }

    fmt.Println("Доступные команды:")
    cmd := []string{}

    for name := range commandsList {
        cmd = append(cmd, name)
    }

    slices.Sort(cmd)
    for _, name := range cmd {
        fmt.Printf("%s\n  %s\n", name, commandsList[name].Description)
    }
    return nil
}
