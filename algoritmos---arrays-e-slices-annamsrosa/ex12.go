package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 7, 9, 12}
	mults := []int{}
	for i := range arr {
		if arr[i]%3 == 0 {
			mults = append(mults, arr[i])
		}
	}
	fmt.Print(mults)
}
