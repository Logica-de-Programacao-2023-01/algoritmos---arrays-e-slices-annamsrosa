package main

import "fmt"

func main() {
	n := []float64{2.1, 3.2, 4.1, 5.4}
	var produto float64 = 1
	for i := range n {
		produto *= n[i]
	}
	fmt.Printf("Produto: %.2f", produto)
}
