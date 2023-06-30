package main

import "fmt"

func primo(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, i, c int = 0, 1, 0
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	n--
	array := []int{}
	for c <= n {
		if primo(i) {
			array = append(array, i)
			c++
		}
		i++
	}
	fmt.Print(array)
}
