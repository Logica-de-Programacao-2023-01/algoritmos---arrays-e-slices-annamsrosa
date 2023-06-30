package main

import "fmt"

func main() {
	n := []int{7, 1, 8, 5, 12, 23}
	var menor, maior int = n[0], n[0]
	for i := range n {
		if n[i] < menor {
			menor = n[i]
		} else if n[i] > maior {
			maior = n[i]
		}
	}
	fmt.Println("Max: ", maior)
	fmt.Print("Min: ", menor)
}
