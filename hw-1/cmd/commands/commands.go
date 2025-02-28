package commands

import (
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
