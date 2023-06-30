package main

import "fmt"

func main() {
	tamanho, soma := 0, 0
	fmt.Print("Digite o tamanho do slice? ")
	fmt.Scan(&tamanho)
	n := make([]int, tamanho)
	for i := range n {
		fmt.Printf("%dº número: \t", i+1)
		fmt.Scan(&n[i])
		soma += n[i]
	}
	fmt.Println()
	for i := range n {
		fmt.Printf("%dº valor: %d\t", i+1, n[i])
	}
	fmt.Printf("\nSoma: %d", soma)

}
