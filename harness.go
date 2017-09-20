package main

import (
	"fmt"
	"github.com/mbwk/octavia/taviserver"
	"os"
)

func main() {
	serverAddress := os.Args[1]
	q, err := taviserver.Query(serverAddress)
	if err != nil {
		fmt.Println("Error encountered, terminating prematurely")
		return
	}
	fmt.Println("Speed (kbps):", q.SpeedBytesPerSec/1024)
}
