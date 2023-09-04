package main

import (
	"fmt"

	api "github.com/JayceFR/library/src/api"
)

func main() {
	fmt.Println("Sup Buddy")
	server := api.NewApiServer(":3000")
	server.Run()
}
