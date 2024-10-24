package stresstest

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RunStressTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	requestChan := make(chan int, totalRequests)
	resultChan := make(chan *http.Response, totalRequests)

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requestChan {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("Falid request:", err)
				} else {
					resultChan <- resp
				}
			}
		}()
	}

	for i := 0; i < totalRequests; i++ {
		requestChan <- i
	}
	close(requestChan)

	wg.Wait()
	close(resultChan)

	totalTime := time.Since(startTime)

	GenerateReport(resultChan, totalRequests, totalTime)
}
