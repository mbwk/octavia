package taviserver

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

type QueryResult struct {
	SpeedBytesPerSec int64
}

func Query(streamAddress string) (*QueryResult, error) {
	resp, err := http.Get(streamAddress)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	bytesBuffer := make([]byte, 4096)

	start := time.Now()
	interval := 10
	intervalDuration := time.Duration(interval) * time.Second
	limit := start.Add(intervalDuration)
	totalRead := 0
	for {
		n, err := reader.Read(bytesBuffer)
		if err != nil {
			break
		}
		totalRead += n
		t := time.Now()
		if t.After(limit) {
			break
		}
	}
	fmt.Println("Read", totalRead, "bytes over", interval, "seconds.")

	if err != nil {
		fmt.Println(err)
	}

	// return success
	result := new(QueryResult)
	result.SpeedBytesPerSec = int64(totalRead / interval)
	return result, nil
}
