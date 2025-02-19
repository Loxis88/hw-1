package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hw-1/handlers"
	"hw-1/services"
	"hw-1/storage"
)

type command string

const (
	AcceptOrderCommand   command = "accept-order"
	ReturnOrderCommand   command = "return-order"
	ProcessOrdersCommand command = "process-orders"
	ListOrdersCommand    command = "list-orders"
	ListReturnsCommand   command = "list-returns"
	OrderHistoryCommand  command = "order-history"
	ImportOrders         command = "import"
	HelpCommand          command = "help"
	ExitCommand          command = "exit"
)

func main() {
	store, err := storage.NewJsonStorage("data.json")
	if err != nil {
		panic(err)
	}

	var service services.OrderServiceInterface = services.New(store)
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

		// Разбиваем ввод на аргументы
		args := strings.Fields(input)
		mainArg := command(args[0])

		// Обновляем os.Args для совместимости с существующими обработчиками
		os.Args = args

		switch mainArg {
		case ExitCommand:
			return

		case HelpCommand:
			fmt.Println(helpMessage)

		case AcceptOrderCommand:
			handlers.HandleAcceptOrder(service)

		case ReturnOrderCommand:
			handlers.HandleReturnOrder(service)

		case ProcessOrdersCommand:
			handlers.HandleProcessOrders(service)

		case ListOrdersCommand:
			handlers.HandleListOrders(service)

		case ListReturnsCommand:
			handlers.HandleListReturns(service)

		case OrderHistoryCommand:
			handlers.HandleOrderHistory(service)

		case ImportOrders:
			handlers.HandleImportOrders(service)
		default:
			fmt.Println("Введите 'help' для списка доступных команд")
		}
	}
}

const helpMessage = `Доступные команды:

1. Принять заказ от курьера
   Команда: accept-order --order-id <ID> --receiver-id <ID> --storage-duration <DAYS>
   Описание: Принимает заказ от курьера и сохраняет его в файл.
   Аргументы:
     --order-id          Идентификатор заказа (обязательный).
     --receiver-id       Идентификатор получателя (обязательный).
     --storage-duration  Срок хранения заказа в днях (обязательный).
   Пример:
     accept-order --order-id 123 --receiver-id 456 --storage-duration 7

2. Вернуть заказ курьеру
   Команда: return-order --order-id <ID>
   Описание: Возвращает заказ курьеру, если срок хранения истек и заказ не был выдан клиенту.
   Аргументы:
     --order-id          Идентификатор заказа (обязательный).
   Пример:
     return-order --order-id 123

3. Выдать заказы и принять возвраты клиента
   Команда: process-orders --client-id <ID> --order-ids <ID1,ID2,...> --action <ACTION>
   Описание: Выдает заказы клиенту или принимает возвраты.
   Аргументы:
     --client-id         Идентификатор клиента (обязательный).
     --order-ids         Список идентификаторов заказов через запятую (обязательный).
     --action            Действие: "issue" (выдать) или "return" (принять возврат) (обязательный).
   Пример:
     process-orders --client-id 456 --order-ids 123,789 --action issue

4. Получить список заказов
   Команда: list-orders --client-id <ID> [--limit <N>]
   Описание: Возвращает список заказов для указанного клиента.
   Аргументы:
     --client-id         Идентификатор клиента (обязательный).
     --limit             Ограничить количество заказов (опционально).
   Пример:
     list-orders --client-id 456 --limit 10 --status in_storage

5. Получить список возвратов
   Команда: list-returns [--page <N>] [--per-page <N>]
   Описание: Возвращает список возвратов с постраничной пагинацией.
   Аргументы:
     --page              Номер страницы (по умолчанию 1).
     --per-page          Количество возвратов на странице (по умолчанию 10).
   Пример:
     list-returns --page 2 --per-page 5

6. Получить историю заказов
   Команда: order-history
   Описание: Возвращает историю заказов в порядке изменения их последнего состояния.
   Аргументы:
        --limit             Ограничить количество заказов (опционально).
   Пример:
     order-history --limit 10

7. Принять заказы из json
    Команда: import
    Описание: Принимает заказы из json файла и сохраняет их в файл.
    Аргументы:
        --path           Путь к файлу (обязательный).

8. Помощь
   Команда: help
   Описание: Выводит список доступных команд и их описание.
   Пример:
     help`
