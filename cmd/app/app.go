package app

import (
	"hw-1/cmd/commands"
	"hw-1/handlers"
	"hw-1/services"
)

func PrepareAndRun(service services.OrderServiceInterface) {
	handlers := map[string]func(services.OrderServiceInterface) error{
		commands.AcceptOrderCommand:   handlers.HandleAcceptOrder,
		commands.ReturnOrderCommand:   handlers.HandleReturnOrder,
		commands.ProcessOrdersCommand: handlers.HandleProcessOrders,
		commands.ListOrdersCommand:    handlers.HandleListOrders,
		commands.ListReturnsCommand:   handlers.HandleListReturns,
		commands.OrderHistoryCommand:  handlers.HandleOrderHistory,
		commands.ImportOrders:         handlers.HandleImportOrders,
		commands.HelpCommand:          handlers.HandleHelp,
		commands.ExitCommand:          handlers.HandleExit,
	}

	commands.Serve(service, handlers)
}
