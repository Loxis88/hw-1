package commands

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"hw-1/handlers"
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

var Commands = map[string]Command{
	AcceptOrderCommand: {
		Description: "Принять заказ от курьера\n  Использование: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>",
		Handle:      handlers.HandleAcceptOrder,
	},
	ReturnOrderCommand: {
		Description: "Вернуть заказ курьеру\n  Использование: return-order --order-id <ID>",
		Handle:      handlers.HandleReturnOrder,
	},
	ProcessOrdersCommand: {
		Description: "Выдать заказы или принять возвраты\n  Использование: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <issue|return>",
		Handle:      handlers.HandleProcessOrders,
	},
	ListOrdersCommand: {
		Description: "Получить список заказов\n  Использование: list-orders --client-id <ID> [--limit <N>]",
		Handle:      handlers.HandleListOrders,
	},
	ListReturnsCommand: {
		Description: "Получить список возвратов\n  Использование: list-returns [--page <N>] [--per-page <N>]",
		Handle:      handlers.HandleListReturns,
	},
	OrderHistoryCommand: {
		Description: "Получить историю заказов\n  Использование: order-history [--client-id <ID>]",
		Handle:      handlers.HandleOrderHistory,
	},
	ImportOrders: {
		Description: "Импортировать заказы\n  Использование: import [--path <путь к json>]",
		Handle:      handlers.HandleImportOrders,
	},
	ExitCommand: {
		Description: "Выйти из программы",
		Handle:      HandleExit,
	},
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

		// замена os.Args на инпут
		os.Args = args
		// тут кароче костыль потому что я не могу полоижть команду помощи в мапу
		if commandName == HelpCommand {
			HandleHelp()
			continue
		}

		cmd, exists := Commands[commandName]
		if !exists {
			fmt.Println("Invalid command:", commandName)
			HandleHelp()
			continue
		}

		// Выполняем обработчик команды
		err = cmd.Handle(service)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// HandleHelp выводит список доступных команд.
func HandleHelp() {
	if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == HelpCommand {
			fmt.Printf("Ввести инормацию по командам (команде)\n  Использование: help [имя команды] (необязательный параметр)\n")
			return
		}
		if command, ok := Commands[os.Args[1]]; ok {
			fmt.Printf("%s\n  %s\n", os.Args[1], command.Description)
		} else {
			fmt.Printf("command %s not found\n", os.Args[1])
		}
		return
	}

	fmt.Println("Доступные команды:")
	commands := []string{}

	for name, _ := range Commands {
		commands = append(commands, name)
	}

	slices.Sort(commands)
	for _, name := range commands {
		fmt.Printf("%s\n  %s\n", name, Commands[name].Description)
	}
}

// HandleExit завершает выполнение программы.
func HandleExit(service services.OrderServiceInterface) error {
	fmt.Println("Выход из программы...")
	os.Exit(0)
	return nil
}
