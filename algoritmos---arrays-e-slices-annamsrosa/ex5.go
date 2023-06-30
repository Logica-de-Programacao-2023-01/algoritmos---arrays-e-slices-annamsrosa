package main

import "fmt"

func main() {
	arr := [3][2]int{}
	for i := range arr {
		for c := range arr[i] {
			fmt.Printf("Digite o valor da %dª linha e %dª coluna: ", i+1, c+1)
			fmt.Scan(&arr[i][c])
		}
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
}
