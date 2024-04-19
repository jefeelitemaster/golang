package main

import "fmt"

func calculate(a float64, b float64) []float64 {
	// your code goes here
	results := make([]float64, 4, 4)
	results[0] = a + b
	results[1] = a - b
	results[2] = a * b
	results[3] = a / b
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
