package main

import "fmt"

func main() {
	arr := []float64{1, 2, 3, 5, 7, 8, 12, 13, 14, 15}
	NA := []float64{}
	for i := range arr {
		if arr[i] > 5 {
			NA = append(NA, arr[i])
		}
	}
	fmt.Print(NA)
}
