/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package stresstest

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	testUrl         string
	testRequests    int
	testConcurrency int
)

// stressTestCmd represents the stressTest command
var startCmd = &cobra.Command{
	Use:   "start --url exemple.com  --requests 100 --concurrency 10",
	Short: "Start stress test",
	Long:  `Starts the stress test based on the parameters passed, which are the application URL (--url), how many requests will be made (--requests), and how many simultaneous threads (--concurrency)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stressTest called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.Flags().StringVar(&testUrl, "url", "", "Url of the application")
	startCmd.Flags().IntVar(&testRequests, "requests", 10, "Number of requets")
	startCmd.Flags().IntVar(&testConcurrency, "concurrency", 1, "Number of simultaneous threads")

	startCmd.MarkFlagRequired("url")

}
