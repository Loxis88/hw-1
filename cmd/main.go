package main

import (
	"hw-1/services"
)

func main() {
	serv := services.New("data.json")

	serv.Test()
}
