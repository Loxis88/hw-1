// hw-1/cmd/registry/command_registry.go
package registry

import (
	"hw-1/cmd/commands"
	"hw-1/handlers"
)

// RegisterCommands возвращает карту всех доступных команд
func RegisterCommands() map[string]commands.Command {
	return map[string]commands.Command{
		commands.AcceptOrderCommand: {
			Description: "Принять заказ от курьера\n  Использование: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>",
			Handle:      handlers.HandleAcceptOrder,
		},
		commands.ReturnOrderCommand: {
			Description: "Вернуть заказ курьеру\n  Использование: return-order --order-id <ID>",
			Handle:      handlers.HandleReturnOrder,
		},
		commands.ExitCommand: {
			Description: "Выйти из программы",
			Handle:      handlers.HandleExit,
		},
		commands.HelpCommand: {
			Description: "Ввести инормацию по командам (команде)\n  Использование: help [имя команды] (необязательный параметр)\n",
			Handle:      handlers.HandleHelp,
		},
		commands.ImportOrders: {
			Description: "Импортировать заказы\n  Использование: import [--path <путь к json>]",
			Handle:      handlers.HandleImportOrders,
		},
		commands.ListOrdersCommand: {
			Description: "Получить список заказов\n  Использование: list-orders --client-id <ID> [--limit <N>]",
			Handle:      handlers.HandleListOrders,
		},
		commands.ListReturnsCommand: {
			Description: "Получить список возвратов\n  Использование: list-returns [--page <N>] [--per-page <N>]",
			Handle:      handlers.HandleListReturns,
		},
		commands.OrderHistoryCommand: {
			Description: "Получить историю заказов\n  Использование: order-history [--limit]",
			Handle:      handlers.HandleOrderHistory,
		},
		commands.ProcessOrdersCommand: {
			Description: "Выдать заказы или принять возвраты\n  Использование: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <issue|return>",
			Handle:      handlers.HandleProcessOrders,
		},
	}
}
