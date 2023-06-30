package main

import "fmt"

func main() {
	n := []int{15, 20, 35, 40, 55}
	valor := 0
	pres := false
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for i := range n {
		if valor == n[i] {
			pres = true
		}
	}
	if !pres {
		n = append(n, valor)
	}
	fmt.Print("Slice: ", n)
}
