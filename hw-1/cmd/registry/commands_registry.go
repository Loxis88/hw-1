// hw-1/cmd/registry/command_registry.go
package registry

import (
	"hw-1/cmd/commands"
	"hw-1/handlers"
	"hw-1/services"
)

// RegisterCommands возвращает карту всех доступных команд
func RegisterCommands() map[string]commands.Command {
	cmds := make(map[string]commands.Command)
	cmds[commands.AcceptOrderCommand] = commands.Command{
		Description: "Принять заказ от курьера\n  Использование: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>",
		Handle:      handlers.HandleAcceptOrder,
	}
	cmds[commands.ReturnOrderCommand] = commands.Command{
		Description: "Вернуть заказ курьеру\n  Использование: return-order --order-id <ID>",
		Handle:      handlers.HandleReturnOrder,
	}
	cmds[commands.ExitCommand] = commands.Command{
		Description: "Выйти из программы",
		Handle:      handlers.HandleExit,
	}
	cmds[commands.ImportOrders] = commands.Command{
		Description: "Импортировать заказы\n  Использование: import [--path <путь к json>]",
		Handle:      handlers.HandleImportOrders,
	}
	cmds[commands.ListOrdersCommand] = commands.Command{
		Description: "Получить список заказов\n  Использование: list-orders --client-id <ID> [--limit <N>]",
		Handle:      handlers.HandleListOrders,
	}
	cmds[commands.ListReturnsCommand] = commands.Command{
		Description: "Получить список возвратов\n  Использование: list-returns [--page <N>] [--per-page <N>]",
		Handle:      handlers.HandleListReturns,
	}
	cmds[commands.OrderHistoryCommand] = commands.Command{
		Description: "Получить историю заказов\n  Использование: order-history [--limit]",
		Handle:      handlers.HandleOrderHistory,
	}
	cmds[commands.ProcessOrdersCommand] = commands.Command{
		Description: "Выдать заказы или принять возвраты\n  Использование: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <issue|return>",
		Handle:      handlers.HandleProcessOrders,
	}
	// тут костыль с замыканием поэтому добавление хелп команды в мапу всегда должно быть в конце
	cmds[commands.HelpCommand] = commands.Command{
		Description: "Ввести инормацию по командам (команде)\n  Использование: help [имя команды] (необязательный параметр)\n",
		Handle: func(service services.OrderServiceInterface) error {
			return handlers.HandleHelp(service, cmds)
		},
	}
	return cmds
}
