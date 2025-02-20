package commands

import (
	"fmt"
	"hw-1/handlers"
	"hw-1/services"
	"os"
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
	Handle      func(services.OrderServiceInterface) err
}

var Commands = map[string]Command{
	AcceptOrderCommand: {
		Description: "Принять заказ от курьера\n   Использование: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>",
		Handle:      handlers.HandleAcceptOrder,
	},
	ReturnOrderCommand: {
		Description: "Вернуть заказ курьеру\n   Использование: return-order --order-id <ID>",
		Handle:      handlers.HandleReturnOrder,
	},
	ProcessOrdersCommand: {
		Description: "Выдать заказы или принять возвраты\n   Использование: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <issue|return>",
		Handle:      handlers.HandleProcessOrders,
	},
	ListOrdersCommand: {
		Description: "Получить список заказов\n   Использование: list-orders --client-id <ID> [--limit <N>]",
		Handle:      handlers.HandleListOrders,
	},
	ListReturnsCommand: {
		Description: "Получить список возвратов\n   Использование: list-returns [--page <N>] [--per-page <N>]",
		Handle:      handlers.HandleListReturns,
	},
	OrderHistoryCommand: {
		Description: "Получить историю заказов\n   Использование: order-history [--client-id <ID>]",
		Handle:      handlers.HandleOrderHistory,
	},
	ImportOrders: {
		Description: "Импортировать заказы\n   Использование: import <детали>",
		Handle:      handlers.HandleImportOrders,
	},
	ExitCommand: {
		Description: "Выйти из программы",
		Handle:      HandleExit,
	},
}

// HandleHelp выводит список доступных команд.
func HandleHelp(service services.OrderServiceInterface) {
	fmt.Println("Доступные команды:")
	commands := []string{}
	for name, _ := range Commands {
		commands = append(commands, name)
	}
	for _, name := range commands {
		fmt.Printf("  %s\n    %s\n", name, Commands[name].Description)
	}
}

// HandleExit завершает выполнение программы.
func HandleExit(service services.OrderServiceInterface) {
	fmt.Println("Выход из программы...")
	os.Exit(0)
}
