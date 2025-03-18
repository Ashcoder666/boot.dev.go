package slic4

import "fmt"

func printReports(messages []string) {
	for _, msg := range messages {
		printCostReport(func(s string) int {
			return len(s) * 2
		}, msg)
	}
}

// don't touch below this line

func Testt(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
