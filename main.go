package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/stats"
	"math-skills/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	inputFile := os.Args[1]
	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: file '%s' not found\n", inputFile)
			return
		} else {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
	}
	defer file.Close()

	// Create a scanner to read the text
	scanner := bufio.NewScanner(file)

	// Slice to store the data
	var data []float64
	var nonNumericValues int

	// REad file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check for missing values
		if line == "" {
			continue
		}

		// Remove commas from the input string
		line = strings.Replace(line, ",", "", -1)

		// Convert strings to float
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			nonNumericValues++
			continue // skip non-numeric values
		}
		data = append(data, value)
	}

	if len(data) == 0 {
		fmt.Printf("No valid data to process! The file '%s' is empty\n", inputFile)
		return
	}

	if nonNumericValues > 0 {
		fmt.Printf("Skipped %d non-numeric value(s).\n", nonNumericValues)
	}

	var overFlowCount int
	var validData []float64

	for _, value := range data {
		intValue, overflow := utils.FloatToInt(value)
		if overflow {
			overFlowCount++
			fmt.Printf("Warning: Potential overflow detected for value: %.2f. Ignoring in statistics.\n", value)
		} else {
			validData = append(validData, float64(intValue))
		}
	}

	// if overFlowCount > 0 {
	// 	fmt.Printf("Skipped %d value(s) due to potential overflow.\n", overFlowCount)
	// }

	mean := stats.Mean(validData)
	median := stats.Median(validData)
	variance := stats.Variance(validData)
	stDev := math.Sqrt(variance)

	intMean, _ := utils.FloatToInt(mean)
	intMedian, _ := utils.FloatToInt(median)
	intVariance, _ := utils.FloatToInt(variance)
	intStDev, _ := utils.FloatToInt(stDev)

	fmt.Printf("Average: %d\n", intMean)
	fmt.Printf("Median: %d\n", intMedian)
	fmt.Printf("Variance: %d\n", intVariance)
	fmt.Printf("Standard Deviation: %d\n", intStDev)

	// Check error
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file\n")
		return
	}
}
