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

var registeredCommands map[string]Command

func GetRegisteredCommands() map[string]Command {
	return registeredCommands
}

// RegisterCommand — функция для регистрации команд

func RegisterCommands(handlers map[string]func(service services.OrderServiceInterface) error) map[string]Command {
	descriptions := map[string]string{
		AcceptOrderCommand:   "Принять заказ от курьера\n  Использование: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>",
		ReturnOrderCommand:   "Вернуть заказ",
		ProcessOrdersCommand: "Обработать заказы",
		ListOrdersCommand:    "Список заказов",
		ListReturnsCommand:   "Список возвратов",
		OrderHistoryCommand:  "История заказов",
		ImportOrders:         "Импортировать заказы",
		HelpCommand:          "Показать справку",
		ExitCommand:          "Выход из программы",
	}

	commands := make(map[string]Command)

	// Регистрируем команды с соответствующими обработчиками
	for cmdName, handler := range handlers {
		if desc, ok := descriptions[cmdName]; ok {
			commands[cmdName] = Command{
				Description: desc,
				Handle:      handler,
			}
		}
	}

	return commands
}

func Serve(service services.OrderServiceInterface) {
	// Подготовим хендлеры для регистрации
	handlers := PrepareHandlers()

	// Получаем команды из метода RegisterCommands
	commands := RegisterCommands(handlers)

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

		cmd, exists := commands[commandName]
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
