package main

import "fmt"

func main() {
	lista := []string{"Marcos", "Daniel", "Pedro", "Bruno", "Rafael", "Isabela", "Marcos", "Paulo"}
	var x string
	var y int
	fmt.Println("A lista de nomes é ", lista)
	fmt.Println("Digite um nome para removê-lo: ")
	fmt.Scan(&x)
	z := y + 1
	for len(lista) > y {
		if x == lista[y] {
			lista = append(lista[:y], lista[z:]...)
		}
		y++
		z++
	}
	fmt.Print("A lista final é: ", lista)
}
