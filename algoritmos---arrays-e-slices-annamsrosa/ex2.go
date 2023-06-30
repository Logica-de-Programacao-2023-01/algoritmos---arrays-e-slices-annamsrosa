package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	for i := range n {
		fmt.Printf("%dº valor: %d\t", i+1, n[i])
	}
	n = append(n[:3], n[4:]...)
	for i := range n {
		fmt.Printf("%dº valor: %d\t", i+1, n[i])
	}
}
