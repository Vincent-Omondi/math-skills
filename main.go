package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := os.Args[1]
	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: file '%s' not found\n", inputFile)
			return
		} else {
			fmt.Printf("error opening file: %W\n", err)
			return
		}
	}
	defer file.Close()

	// Create a scanner to read the text
	scanner := bufio.NewScanner(file)

	// Slice to store the data
	var data []float64

	// REad file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Convert strings to float
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			// fmt.Printf("Error convering value: %s\n", err)
			continue
		}
		data = append(data, value)
	}

	// data = BubbleSort(data)
	mean := Mean(data)
	fmt.Printf("Average: %.2f\n", mean)
	median := Median(data)
	fmt.Printf("Median: %.2f\n", median)
	variance := Variance(data)
	fmt.Printf("Variance: %.2f\n", variance)
	stdev := StDev(variance)
	fmt.Printf("Standard Deviation: %.2f\n", stdev)

	// Check error
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file\n")
		return
	}
}
