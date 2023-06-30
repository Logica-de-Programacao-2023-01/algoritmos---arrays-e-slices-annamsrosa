package main

import "fmt"

func main() {
	n := [3]int{3, 20, 48}
	valor := 0
	pres := false
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for i := range n {
		if valor == n[i] {
			fmt.Printf("Valor presente na %dª posição.", i+1)
			pres = true
		}
	}
	if pres == false {
		fmt.Println("Número não encontrado")
	}
}
