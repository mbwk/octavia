package main

import (
	"fmt"
	"github.com/mbwk/octavia/taviserver"
	"os"
	"sync"
)

func testSpeed(serverAddress string) {
	q, err := taviserver.Query(serverAddress)
	if err != nil {
		fmt.Println(serverAddress, "-", err)
		return
	}
	fmt.Println(serverAddress, "-", "Speed (kbps):", q.SpeedBytesPerSec/1024)
}

func main() {
	args := os.Args[1:]
	argc := len(args)
	if argc == 0 {
		fmt.Println("No arguments provided, exiting early")
		return
	}

	var wg sync.WaitGroup
	wg.Add(argc)
	for _, url := range args {
		go func(url string) {
			defer wg.Done()
			testSpeed(url)
		}(url)
	}
	wg.Wait()
}
