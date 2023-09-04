package main

import "fmt"

func main() {
	fmt.Println("Sup Buddy")
	server := NewApiServer(":3000")
	server.Run()
}
