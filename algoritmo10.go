package main

import "fmt"

func main() {
	lista := []int{5, 2, 8, 90, 33, 71, -7, 50, -5, 29}
	var maior int = 0
	var menor int = 100
	var n int = 0
	for n < len(lista) {
		if lista[n] > maior {
			maior = lista[n]
		} else if lista[n] < menor {
			menor = lista[n]
		}
		n++
	}
	fmt.Print("O menor número da lista é ", menor, " e o maior número da lista é ", maior)
}
