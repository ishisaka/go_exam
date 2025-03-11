package main

import (
	"fmt"

	"example.com/module"
	"example.com/server"
)

func main() {
	fmt.Println(module.Hello())
	server.StartServer()
}
