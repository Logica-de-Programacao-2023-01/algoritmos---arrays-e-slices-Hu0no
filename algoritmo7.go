package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	var n int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&n)
	if n >= 1 && n <= 5 {
		fmt.Print("A lista resultante é ", lista)
	} else {
		lista = append(lista, n)
		fmt.Print("A lista resultante é ", lista)
	}
}
