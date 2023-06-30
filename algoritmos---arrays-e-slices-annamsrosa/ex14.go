package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var ind1, ind2, aux int
	fmt.Print("Indice 1: ")
	fmt.Scan(&ind1)
	fmt.Print("Indice 2: ")
	fmt.Scan(&ind2)
	aux = arr[ind1]
	arr[ind1] = arr[ind2]
	arr[ind2] = aux
	fmt.Print(arr)

}
