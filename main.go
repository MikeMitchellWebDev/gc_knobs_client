package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	path       = flag.String("p", "", "string")
	goroutines = flag.Int("g", 4, "int")
	sleep      = flag.Int("s", 1, "int")
	repeats    = flag.Int("r", 1, "int")
)

func main() {
	flag.Parse()

	repeat_requests := strconv.Itoa(*repeats)
	sleep_between_requests := strconv.Itoa(*sleep)
	values := map[string]string{"repeat": repeat_requests, "sleep": sleep_between_requests, "path": *path}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println(err)
	}
	wg1 := &sync.WaitGroup{}
	wg1.Add(*goroutines)
	startTime := time.Now()
	for i := 0; i < *goroutines; i++ {
		go func() {
			_, err := http.Post("http://localhost:8000/git_repo", "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Println(err)
			}

			wg1.Done()
		}()

	}
	wg1.Wait()
	stopTime := time.Now()
	executionTime := stopTime.Sub(startTime)
	fmt.Println("execution time", executionTime)
}
