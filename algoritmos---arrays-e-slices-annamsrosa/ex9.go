package main

import "fmt"

func main() {
	arr := [6]float64{}
	fmt.Print("Digite um número: ")
	fmt.Scan(&arr[0])
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[0]
	}
	fmt.Print("Array: ", arr)
}
