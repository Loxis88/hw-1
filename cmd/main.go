package main

import (
	"flag"
	"fmt"
	"os"
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
	HelpCommand          command = "help"
)

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
   Команда: list-orders --client-id <ID> [--limit <N>] [--status <STATUS>]
   Описание: Возвращает список заказов для указанного клиента.
   Аргументы:
     --client-id         Идентификатор клиента (обязательный).
     --limit             Ограничить количество заказов (опционально).
     --status            Фильтр по статусу заказа (например, "in_storage") (опционально).
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
   Команда: order-history --client-id <ID>
   Описание: Возвращает историю заказов в порядке изменения их последнего состояния.
   Аргументы:
     --client-id         Идентификатор клиента (обязательный).
   Пример:
     order-history --client-id 456

7. Помощь
   Команда: help
   Описание: Выводит список доступных команд и их описание.
   Пример:
     help`

func main() {
	store, err := storage.NewJsonStorage("data.json")
	if err != nil {
		panic(err)
	}
	var service services.OrderServiceInterface = services.New(store)
	_ = service

	mainArg := command(os.Args[1])

	switch mainArg {
	case HelpCommand:
		fmt.Println(helpMessage)
		return

	case AcceptOrderCommand:
		handlers.HandleAcceptOrder(service)
		return

	case ReturnOrderCommand:
		handlers.HandleReturnOrder(service)
		return

	case ProcessOrdersCommand:
		handlers.HandleProcessOrders(service)
		return

	case ListOrdersCommand:
		handlers.HandleListOrders(service)
		return

	case ListReturnsCommand:
		handlers.HandleListReturns(service)
		return

	case OrderHistoryCommand:
		handlers.HandleOrderHistory(service)
		return

	default:
		fmt.Println("Invalid command")
		fmt.Println(helpMessage)
		return
	}
}
