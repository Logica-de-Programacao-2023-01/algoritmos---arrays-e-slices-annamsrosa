package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	soma := []int{}
	for i := range arr1 {
		soma = append(soma, arr1[i])
	}
	for i := range arr2 {
		soma = append(soma, arr2[i])
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(soma)
}
