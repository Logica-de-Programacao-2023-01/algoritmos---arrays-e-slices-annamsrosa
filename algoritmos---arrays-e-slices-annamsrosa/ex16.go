package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	NA := []int{}
	for i := range arr {
		if arr[i]%2 == 0 {
			NA = append(NA, arr[i])
		}
	}
	fmt.Print(NA)
}
