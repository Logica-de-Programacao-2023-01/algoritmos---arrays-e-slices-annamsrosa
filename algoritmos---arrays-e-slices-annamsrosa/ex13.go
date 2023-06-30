package main

import "fmt"

func main() {
	arr := [7]int{}
	fmt.Print("Primeiro elemento do array: ")
	fmt.Scan(&arr[0])
	fmt.Print("Último elemento do array: ")
	fmt.Scan(&arr[6])
	for i := 1; i < len(arr)-1; i++ {
		arr[i] = arr[0] + i
	}
	fmt.Print(arr)
}
