package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hw-1/services"
)

const (
	AcceptOrderCommand   = "accept-order"
	ReturnOrderCommand   = "return-order"
	ProcessOrdersCommand = "process-orders"
	ListOrdersCommand    = "list-orders"
	ListReturnsCommand   = "list-returns"
	OrderHistoryCommand  = "order-history"
	ImportOrders         = "import"
	HelpCommand          = "help"
	ExitCommand          = "exit"
)

type Command struct {
	Description string
	Handle      func(services.OrderServiceInterface) error
}

var Commands = make(map[string]Command)

// RegisterCommand — функция для регистрации команд
func RegisterCommand(name string, cmd Command) {
	Commands[name] = cmd
}

// наверное стоит это вынести куда то в другое место но я пока не знаю куда
func Serve(service services.OrderServiceInterface) {
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

		os.Args = args

		cmd, exists := Commands[commandName]
		if !exists {
			fmt.Println("Invalid command:", commandName, "type help for more information")
			continue
		}

		// Выполняем обработчик команды
		err = cmd.Handle(service)
		if err != nil {
			fmt.Println(err)
		}
	}
}
