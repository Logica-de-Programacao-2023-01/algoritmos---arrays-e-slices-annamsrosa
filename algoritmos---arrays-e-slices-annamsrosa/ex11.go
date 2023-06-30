package main

import "fmt"

func main() {
	matriz := [2][3]int{{5, 15, 25}, {10, 20, 30}}
	var indicel, indicec int
	fmt.Println(matriz)
linha:
	fmt.Print("Índice de linha:")
	fmt.Scan(&indicel)
	if indicel > 1 {
		fmt.Println("Número invalido 1")
		goto linha
	}

coluna:
	fmt.Print("Índice de coluna:")
	fmt.Scan(&indicec)

	if indicec > 2 {
		fmt.Println("Número invalido ")
		goto coluna
	}

	fmt.Println(matriz[indicel][indicec])

}
