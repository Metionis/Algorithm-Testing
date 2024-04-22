package main

import "fmt"

func main() {
	// Stock prices
	stockPrices := []int{2, 3, 5, 6}
	// Stock names
	stockNames := []string{"APPL", "IBM", "TATA"}
	// Stock data
	stockData := []map[string]interface{}{
		{"ticker": "APPL", "price": 302},
		{"ticker": "IBM", "price": 312},
		{"ticker": "TATA", "price": 354},
	}

	// Dimension Array
	stockPrizes := [][]int{
		{2, 3, 5, 6},
		{12, 14, 15, 16},
		{5, 7, 1, 9},
	}

	// Printing stock prices
	fmt.Println("stockPrices:")
	for _, price := range stockPrices {
		fmt.Println(price)
	}

	// Printing stock names
	fmt.Println("\nstockNames:")
	for _, name := range stockNames {
		fmt.Println(name)
	}

	// Printing stock data
	fmt.Println("\nstockData:")
	for _, data := range stockData {
		fmt.Printf("Ticker: %s, Price: %d\n", data["ticker"], data["price"])
	}

	// Printing stock prizes
	fmt.Println("\nstockPrizes:")
	for _, row := range stockPrizes {
		for _, prize := range row {
			fmt.Printf("%d ", prize)
		}
		fmt.Println() // Print new line for each row
	}
}
