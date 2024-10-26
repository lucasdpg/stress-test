package stresstest

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RunStressTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	var successes, failures int
	var mu sync.Mutex
	statusDistribution := make(map[int]int)
	startTime := time.Now()

	sendRequests := func() {
		defer wg.Done()
		for i := 0; i < totalRequests; i++ {
			resp, err := http.Get(url)
			mu.Lock()

			if err != nil {
				failures++
			} else {
				statusDistribution[resp.StatusCode]++
				if resp.StatusCode == 200 {
					successes++
				} else {
					failures++
				}
				resp.Body.Close()
			}

			mu.Unlock()
		}
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go sendRequests()
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	fmt.Printf("\nStress Test Completed\n")
	fmt.Printf("Total requests: %d\n", totalRequests*concurrency)
	fmt.Printf("Total time: %v\n", totalTime)
	fmt.Printf("Total successes: %d\n", successes)
	fmt.Printf("Total failures: %d\n", failures)
	fmt.Printf("HTTP status code distribution:\n")
	for code, count := range statusDistribution {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
