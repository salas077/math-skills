package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := ReadData(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Println("Error: file is empty")
		os.Exit(1)
	}

	average := Average(data)
	median := Median(data)
	variance := Variance(data)
	stdDev := StandardDeviation(data)

	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}

func ReadData(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func Average(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range data {
		sum = sum + num
	}
	return sum / float64(len(data))
}

func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	n := len(data)
	sorted := make([]float64, n)

	for i := 0; i < n; i++ {
		sorted[i] = data[i]
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				temp := sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = temp
			}
		}
	}

	if n%2 == 1 {
		return sorted[n/2]
	} else {
		left := sorted[n/2-1]
		right := sorted[n/2]
		return (left + right) / 2.0
	}
}

func Variance(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	avg := Average(data)
	sumSquaredDiff := 0.0

	for _, num := range data {
		diff := num - avg
		sumSquaredDiff += diff * diff
	}

	return sumSquaredDiff / float64(len(data))
}

func StandardDeviation(data []float64) float64 {
	variance := Variance(data)
	return math.Sqrt(variance)
}
