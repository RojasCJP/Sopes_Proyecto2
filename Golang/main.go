package main

import (
	"fmt"
	"servidor/gRPC"
	"servidor/gRPC/server"
)

func main() {
	fmt.Println("hola buenas")
	go server.Export()
	gRPC.LevantarServidor()
}
