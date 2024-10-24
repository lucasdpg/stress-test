package stresstest

import (
	"fmt"
	"net/http"
	"time"
)

func GenerateReport(results chan *http.Response, totalRequests int, totalTime time.Duration) {
	var successCount, failCount int
	statusCodeCount := make(map[int]int)

	for resp := range results {
		statusCodeCount[resp.StatusCode]++
		if resp.StatusCode == 200 {
			successCount++
		} else {
			failCount++
		}
	}

	fmt.Printf("Stress Test Completed\n")
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Total time: %s\n", totalTime)
	fmt.Printf("Successes: %d\n", successCount)
	fmt.Printf("Failures: %d\n", failCount)
	fmt.Println("HTTP status code distribution:")
	for code, count := range statusCodeCount {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
