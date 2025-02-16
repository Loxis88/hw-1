package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
		orderID := flag.Uint("order-id", 0, "orderID")
		receiverID := flag.Uint("receiver-id", 0, "receiverID")
		storageDuration := flag.Uint("storage-duration", 0, "duration")

		flag.Parse()

		if flag.NFlag() != 3 || *orderID == 0 || *receiverID == 0 || *storageDuration == 0 {
			fmt.Println("Invalid arguments")
			return
		}

		if err := service.AcceptOrder(*orderID, *receiverID, time.Now().Add(time.Duration(*storageDuration)*time.Hour*24)); err != nil {
			fmt.Println("Error accepting order:", err)
		}

		fmt.Println("Order accepted successfully")
		return

	case ReturnOrderCommand:
		var orderID = flag.Uint("order-id", 0, "orderID")
		flag.Parse()

		if flag.NFlag() != 1 || *orderID == 0 {
			fmt.Println("Invalid arguments")
			return
		}

		if err := service.ReturnOrderToCourier(*orderID); err != nil {
			fmt.Println("Error returning order:", err)
		}

		fmt.Println("Order returned successfully")
		return

	case ProcessOrdersCommand:
		clientID := flag.Uint("client-id", 0, "clientID")
		orderIDs := flag.String("order-ids", "", "orderIDs")
		action := flag.String("action", "", "action")
		flag.Parse()

		if flag.NFlag() < 3 {
			fmt.Println("Invalid arguments")
			return
		}
		if *action != "return" && *action != "issue" {
			fmt.Println("Invalid action")
			return
		}

		orders := strings.Split(*orderIDs, ",")
		var ids []uint = make([]uint, len(orders))

		for i, _ := range orders {
			id, err := strconv.Atoi(orders[i])
			if err != nil {
				fmt.Println("Invalid order ID:", orders[i])
				return
			}
			ids[i] = uint(id)
		}
		switch *action {
		case "return":
			if err := service.AcceptReturns(*clientID, ids...); err != nil {
				fmt.Println("Error accepting orders:", err)
				return
			}
			fmt.Println("Заказы успешно приняты")
			return
		case "issue":
			if err := service.DeliverOrders(*clientID, ids...); err != nil {
				fmt.Println("Error returning orders:", err)
				return
			}
			fmt.Println("Заказы успешно выданы")
			return
		}
		return
	case ListOrdersCommand:
		customerID := flag.Uint("client-id", 0, "clientID")
		limit := flag.Int("limit", 0, "limit")
		flag.Parse()

		if *customerID == 0 {
			fmt.Println("Invalid client ID")
			return
		}

		orders, err := service.GetCustomerOrders(*customerID, *limit)
		if err != nil {
			fmt.Println("Error listing orders:", err)
			return
		}
		fmt.Println("Orders:", orders)
		return

	case ListReturnsCommand:
		returns, err := service.GetReturnedOrders()
		if err != nil {
			fmt.Println("Error listing returns:", err)
			return
		}
		fmt.Println("Returns:", returns)
		return

	case OrderHistoryCommand:
		customerID := flag.Int("client-id", 0, "clientID")
		flag.Parse()

		if *customerID == 0 {
			fmt.Println("Invalid client ID")
			return
		}

		history, err := service.GetOrderHistory(*customerID)
		if err != nil {
			fmt.Println("Error listing order history:", err)
			return
		}
		fmt.Println("Order History:", history)
		return

	default:
		fmt.Println("Invalid command")
		fmt.Println(helpMessage)
		return
	}
}

/*
service.AcceptOrder(10, 15, time.Now().Add(time.Hour))

	// Call GetOrderHistory method and print the result
	orderHistory, err := service.GetOrderHistory(10)
	if err != nil {
		fmt.Println("Error fetching order history:", err)
	} else {
		fmt.Println("Order History:", orderHistory)
	}

	// Call GetReturnedOrders method and print the result
	returnedOrders, err := service.GetReturnedOrders()
	if err != nil {
		fmt.Println("Error fetching returned orders:", err)
	} else {
		fmt.Println("Returned Orders:", returnedOrders)
	}
*/
