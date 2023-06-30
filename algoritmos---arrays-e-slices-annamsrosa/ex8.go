package main

import "fmt"

func main() {
	slice := make([]string, 8)

	slice[0] = "Ana"
	slice[1] = "Bianca"
	slice[2] = "Carlos"
	slice[3] = "Daniel"
	slice[4] = "Estela"
	slice[5] = "Fabiana"
	slice[6] = "Gabriel"
	slice[7] = "Gabriel"

	var value string
	fmt.Print("Digite um valor: ")
	fmt.Scan(&value)

	result := make([]string, 0)

	for _, item := range slice {
		if item != value {
			result = append(result, item)
		}
	}

	fmt.Println("Slice resultante:", result)
}
