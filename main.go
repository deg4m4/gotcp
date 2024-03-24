package main

import (
	"deg4m4/gotcp/client"
	"deg4m4/gotcp/server"
	"fmt"
)

func main() {
	var choice string

	fmt.Print("Enter 'server' or 'client': ")
	fmt.Scanln(&choice)

	switch choice {
	case "server", "s":
		server.ServerRun()
	case "client", "c":
		client.ClientRun()
	default:
		fmt.Println("Invalid choice")
	}
}


