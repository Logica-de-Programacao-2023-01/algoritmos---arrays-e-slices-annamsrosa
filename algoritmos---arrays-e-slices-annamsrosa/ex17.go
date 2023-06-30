package main

import "fmt"

func main() {
	arr := []int{5, 2, 5, 4, 5, 6, 5, 8, 5, 10}
	soma := 0
	for i := range arr {
		if i%2 == 0 {
			soma += arr[i]
		}
	}
	fmt.Print(soma)
}
