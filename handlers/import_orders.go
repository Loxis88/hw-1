package handlers

import (
	"flag"
	"fmt"
	"os"

	"hw-1/services"
)

func HandleImportOrders(service services.OrderServiceInterface) {
	flagSet := flag.NewFlagSet("path", flag.ExitOnError)

	path := flagSet.String("path", "", "path")

	flagSet.Parse(os.Args)

	fmt.Println(*path)
	if *path == "" {
		fmt.Println("error while parsin arguments")
	}

	err := service.ImportOrders(*path)
	if err != nil {
		fmt.Print(err)
	}
}
