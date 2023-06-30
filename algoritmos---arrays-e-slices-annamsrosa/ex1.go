package main

import "fmt"

func main() {
	n := []int{1, 2, 3}
	soma := 0
	for i := range n {
		soma += n[i]
		fmt.Println("Valor: ", i+1, n[i])
	}
	fmt.Println("\nSoma: ", soma)
}
