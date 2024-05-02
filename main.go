package main

import (
	"library/server"
)

func main() {

	server := server.Build()
	server.Run()
}
