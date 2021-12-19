package main

import (
	"fmt"
	"servidor/gRPC/client"
	"servidor/gRPC/server"
)

func main() {
	fmt.Println("hola buenas")
	go client.Export()
	server.Export()
}
