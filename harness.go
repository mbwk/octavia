package main

import (
	"fmt"
	"github.com/mbwk/octavia/taviserver"
	"os"
)

func testSpeed(serverAddress string) string {
	q, err := taviserver.Query(serverAddress)
	if err != nil {
		return fmt.Sprint(serverAddress, " - ", err)
	}
	return fmt.Sprint(serverAddress, " - Speed (kbps): ", q.SpeedBytesPerSec/1024)
}

func main() {
	args := os.Args[1:]
	argc := len(args)
	if argc == 0 {
		fmt.Println("No arguments provided, exiting early")
		return
	}

	items := len(args)
	c := make(chan string)
	for _, url := range args {
		go func(url string, c chan string) {
			c <- testSpeed(url)
		}(url, c)
	}

	for i := 0; i < items; i++ {
		fmt.Println(<-c)
	}
}
