package cli

import (
	"fmt"

	stresstest "github.com/lucasdpg/stress-test/pkg/stressTest"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the stress test",
	Run: func(cmd *cobra.Command, args []string) {
		if requests <= 0 {
			fmt.Println("The number of requests must be greater than zero.")
			return
		}
		if concurrency <= 0 {
			fmt.Println("The concurrency must be greater than zero.")
			return
		}
		stresstest.RunStressTest(url, requests, concurrency)
	},
}

func init() {
	startCmd.Flags().StringVar(&url, "url", "", "Service URL (required)")
	startCmd.Flags().IntVar(&requests, "requests", 10, "Total number of requests")
	startCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Number of concurrent requests")
	startCmd.MarkFlagRequired("url")
}
