package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 18, 2}
	crescente := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			crescente = false
			break
		}
	}
	if len(arr) == 1 {
		crescente = true
	}
	if crescente {
		fmt.Println("Array está em ordem crescente")
	} else {
		fmt.Println("Array está desordenado")
	}
}
